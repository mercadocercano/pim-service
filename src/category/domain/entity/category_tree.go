package entity

import (
	"errors"
	"sort"
)

// ErrInvalidCategoryTree representa errores relacionados con el árbol de categorías
var ErrInvalidCategoryTree = errors.New("árbol de categorías inválido")

// CategoryTreeNode representa un nodo en el árbol de categorías
type CategoryTreeNode struct {
	Category      *MarketplaceCategory
	Children      []*CategoryTreeNode
	Parent        *CategoryTreeNode
	Depth         int
	Path          []string // Slugs desde la raíz hasta este nodo
	FullPath      string   // Path completo como string "parent/child/grandchild"
	ChildrenCount int      // Total de descendientes
}

// CategoryTree representa el árbol completo de categorías marketplace
type CategoryTree struct {
	RootNodes  []*CategoryTreeNode
	AllNodes   map[string]*CategoryTreeNode // Mapa por ID para acceso rápido
	MaxDepth   int
	TotalNodes int
}

// NewCategoryTree crea un nuevo árbol de categorías a partir de una lista plana
func NewCategoryTree(categories []MarketplaceCategory) (*CategoryTree, error) {
	if len(categories) == 0 {
		return &CategoryTree{
			RootNodes:  []*CategoryTreeNode{},
			AllNodes:   make(map[string]*CategoryTreeNode),
			MaxDepth:   0,
			TotalNodes: 0,
		}, nil
	}

	tree := &CategoryTree{
		RootNodes: []*CategoryTreeNode{},
		AllNodes:  make(map[string]*CategoryTreeNode),
		MaxDepth:  0,
	}

	// Crear nodos para todas las categorías
	for i := range categories {
		category := &categories[i]
		node := &CategoryTreeNode{
			Category:      category,
			Children:      []*CategoryTreeNode{},
			Depth:         category.Level,
			Path:          []string{},
			ChildrenCount: 0,
		}
		tree.AllNodes[category.ID] = node
	}

	// Construir relaciones padre-hijo
	for _, node := range tree.AllNodes {
		if node.Category.ParentID == nil {
			// Es un nodo raíz
			tree.RootNodes = append(tree.RootNodes, node)
		} else {
			// Buscar el padre y agregarse como hijo
			parent, exists := tree.AllNodes[*node.Category.ParentID]
			if !exists {
				return nil, errors.New("categoría padre no encontrada: " + *node.Category.ParentID)
			}
			parent.Children = append(parent.Children, node)
			node.Parent = parent
		}
	}

	// Calcular paths y métricas
	tree.calculatePaths()
	tree.calculateChildrenCounts()
	tree.sortNodes()
	tree.TotalNodes = len(categories)

	return tree, nil
}

// GetNodeByID retorna un nodo por su ID
func (ct *CategoryTree) GetNodeByID(categoryID string) (*CategoryTreeNode, error) {
	node, exists := ct.AllNodes[categoryID]
	if !exists {
		return nil, errors.New("categoría no encontrada: " + categoryID)
	}
	return node, nil
}

// GetNodeBySlug retorna un nodo por su slug
func (ct *CategoryTree) GetNodeBySlug(slug string) (*CategoryTreeNode, error) {
	for _, node := range ct.AllNodes {
		if node.Category.Slug == slug {
			return node, nil
		}
	}
	return nil, errors.New("categoría no encontrada con slug: " + slug)
}

// GetRootCategories retorna todas las categorías raíz ordenadas
func (ct *CategoryTree) GetRootCategories() []*MarketplaceCategory {
	categories := make([]*MarketplaceCategory, len(ct.RootNodes))
	for i, node := range ct.RootNodes {
		categories[i] = node.Category
	}
	return categories
}

// GetCategoriesByLevel retorna todas las categorías de un nivel específico
func (ct *CategoryTree) GetCategoriesByLevel(level int) []*MarketplaceCategory {
	var categories []*MarketplaceCategory
	for _, node := range ct.AllNodes {
		if node.Category.Level == level && node.Category.IsActive {
			categories = append(categories, node.Category)
		}
	}

	// Ordenar por SortOrder
	sort.Slice(categories, func(i, j int) bool {
		return categories[i].SortOrder < categories[j].SortOrder
	})

	return categories
}

// GetChildren retorna los hijos directos de una categoría
func (ct *CategoryTree) GetChildren(categoryID string) ([]*MarketplaceCategory, error) {
	node, err := ct.GetNodeByID(categoryID)
	if err != nil {
		return nil, err
	}

	children := make([]*MarketplaceCategory, len(node.Children))
	for i, child := range node.Children {
		children[i] = child.Category
	}
	return children, nil
}

// GetAncestors retorna todos los ancestros de una categoría (desde raíz hasta padre)
func (ct *CategoryTree) GetAncestors(categoryID string) ([]*MarketplaceCategory, error) {
	node, err := ct.GetNodeByID(categoryID)
	if err != nil {
		return nil, err
	}

	var ancestors []*MarketplaceCategory
	current := node.Parent
	for current != nil {
		ancestors = append([]*MarketplaceCategory{current.Category}, ancestors...)
		current = current.Parent
	}
	return ancestors, nil
}

// GetDescendants retorna todos los descendientes de una categoría
func (ct *CategoryTree) GetDescendants(categoryID string) ([]*MarketplaceCategory, error) {
	node, err := ct.GetNodeByID(categoryID)
	if err != nil {
		return nil, err
	}

	var descendants []*MarketplaceCategory
	ct.collectDescendants(node, &descendants)
	return descendants, nil
}

// GetFullPath retorna el path completo de una categoría
func (ct *CategoryTree) GetFullPath(categoryID string) (string, error) {
	node, err := ct.GetNodeByID(categoryID)
	if err != nil {
		return "", err
	}
	return node.FullPath, nil
}

// GetBreadcrumbs retorna la ruta de navegación hasta una categoría
func (ct *CategoryTree) GetBreadcrumbs(categoryID string) ([]MarketplaceCategory, error) {
	node, err := ct.GetNodeByID(categoryID)
	if err != nil {
		return nil, err
	}

	breadcrumbs := make([]MarketplaceCategory, len(node.Path))
	for i, slug := range node.Path {
		breadcrumbNode, err := ct.GetNodeBySlug(slug)
		if err != nil {
			return nil, err
		}
		breadcrumbs[i] = *breadcrumbNode.Category
	}
	return breadcrumbs, nil
}

// ValidateHierarchy verifica que la jerarquía sea válida
func (ct *CategoryTree) ValidateHierarchy() error {
	// Verificar que no hay ciclos
	for _, node := range ct.AllNodes {
		if ct.hasCycle(node) {
			return errors.New("se detectó un ciclo en la jerarquía de categorías")
		}
	}

	// Verificar que la profundidad no excede el máximo
	if ct.MaxDepth > 3 {
		return errors.New("la jerarquía excede la profundidad máxima de 3 niveles")
	}

	return nil
}

// ToFlatList convierte el árbol a una lista plana ordenada jerárquicamente
func (ct *CategoryTree) ToFlatList() []*MarketplaceCategory {
	var flatList []*MarketplaceCategory
	for _, rootNode := range ct.RootNodes {
		ct.collectInOrder(rootNode, &flatList)
	}
	return flatList
}

// Métodos privados

// calculatePaths calcula los paths y full paths para todos los nodos
func (ct *CategoryTree) calculatePaths() {
	for _, rootNode := range ct.RootNodes {
		ct.calculateNodePath(rootNode, []string{})
	}
}

// calculateNodePath calcula el path de un nodo específico
func (ct *CategoryTree) calculateNodePath(node *CategoryTreeNode, parentPath []string) {
	// Construir el path actual
	node.Path = append(parentPath, node.Category.Slug)
	node.FullPath = ""
	for i, slug := range node.Path {
		if i > 0 {
			node.FullPath += "/"
		}
		node.FullPath += slug
	}

	// Actualizar MaxDepth
	if node.Depth > ct.MaxDepth {
		ct.MaxDepth = node.Depth
	}

	// Procesar hijos
	for _, child := range node.Children {
		ct.calculateNodePath(child, node.Path)
	}
}

// calculateChildrenCounts calcula el conteo de descendientes para cada nodo
func (ct *CategoryTree) calculateChildrenCounts() {
	for _, rootNode := range ct.RootNodes {
		ct.calculateNodeChildrenCount(rootNode)
	}
}

// calculateNodeChildrenCount calcula el conteo de descendientes de un nodo
func (ct *CategoryTree) calculateNodeChildrenCount(node *CategoryTreeNode) int {
	count := len(node.Children)
	for _, child := range node.Children {
		count += ct.calculateNodeChildrenCount(child)
	}
	node.ChildrenCount = count
	return count
}

// sortNodes ordena los nodos por SortOrder en cada nivel
func (ct *CategoryTree) sortNodes() {
	// Ordenar nodos raíz
	sort.Slice(ct.RootNodes, func(i, j int) bool {
		return ct.RootNodes[i].Category.SortOrder < ct.RootNodes[j].Category.SortOrder
	})

	// Ordenar hijos de cada nodo recursivamente
	for _, rootNode := range ct.RootNodes {
		ct.sortNodeChildren(rootNode)
	}
}

// sortNodeChildren ordena los hijos de un nodo recursivamente
func (ct *CategoryTree) sortNodeChildren(node *CategoryTreeNode) {
	sort.Slice(node.Children, func(i, j int) bool {
		return node.Children[i].Category.SortOrder < node.Children[j].Category.SortOrder
	})

	for _, child := range node.Children {
		ct.sortNodeChildren(child)
	}
}

// collectDescendants recolecta todos los descendientes de un nodo
func (ct *CategoryTree) collectDescendants(node *CategoryTreeNode, descendants *[]*MarketplaceCategory) {
	for _, child := range node.Children {
		*descendants = append(*descendants, child.Category)
		ct.collectDescendants(child, descendants)
	}
}

// collectInOrder recolecta nodos en orden jerárquico
func (ct *CategoryTree) collectInOrder(node *CategoryTreeNode, list *[]*MarketplaceCategory) {
	*list = append(*list, node.Category)
	for _, child := range node.Children {
		ct.collectInOrder(child, list)
	}
}

// hasCycle detecta ciclos en la jerarquía
func (ct *CategoryTree) hasCycle(node *CategoryTreeNode) bool {
	visited := make(map[string]bool)
	return ct.detectCycle(node, visited)
}

// detectCycle detecta ciclos recursivamente
func (ct *CategoryTree) detectCycle(node *CategoryTreeNode, visited map[string]bool) bool {
	if visited[node.Category.ID] {
		return true
	}

	visited[node.Category.ID] = true
	for _, child := range node.Children {
		if ct.detectCycle(child, visited) {
			return true
		}
	}
	delete(visited, node.Category.ID)
	return false
}
