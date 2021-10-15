package Q236_Lowest_Common_Ancestor_of_a_Binary_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//lowestCommonAncestor
//先遍历出所有子节点对应的父节点，
//分别对P、Q节点找上级节点并保存至相应的map后，分别在另一个map判断是否存在，存在则返回节点
//Runtime: 25 ms, faster than 13.76% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
//Memory Usage: 7.7 MB, less than 93.12% of Go online submissions for Lowest Common Ancestor of a Binary Tree.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//pp := make([]int, 0)
	mapPath, mapP, mapQ := make(map[*TreeNode]*TreeNode), make(map[*TreeNode]struct{}), make(map[*TreeNode]struct{})
	path, rootTmp := []*TreeNode{root}, root
	// 遍历二叉树
	for len(path) != 0 || rootTmp != nil {
		if rootTmp != nil {
			if rootTmp.Left != nil {
				mapPath[rootTmp.Left] = rootTmp
			}
			if rootTmp.Right != nil {
				mapPath[rootTmp.Right] = rootTmp
			}

			//pp = append(pp, rootTmp.Val)
			if rootTmp.Left != nil {
				path = append(path, rootTmp.Left)
			}
			rootTmp = rootTmp.Left
			continue
		}
		if len(path) != 0 {
			right := path[len(path)-1].Right
			rootTmp = right
			if right != nil {
				path[len(path)-1] = right
			} else {
				path = path[:len(path)-1]
			}
		}
	}
	//for k, v := range mapPath {
	//	fmt.Println("k,v:", k.Val, "|", v.Val)
	//}

	//fmt.Println(pp)
	// 上移
	for {
		mapP[p], mapQ[q] = struct{}{}, struct{}{}
		if _, ok := mapQ[p]; ok {
			return p
		}
		if _, ok := mapP[q]; ok {
			return q
		}
		p, q = mapPath[p], mapPath[q]
	}
}
