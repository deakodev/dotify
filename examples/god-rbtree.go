package main

// tree := redblacktree.NewWithIntComparator()

// tree.Put(1, "x") // 1->x
// tree.Put(2, "b") // 1->x, 2->b (in order)
// tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
// tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
// tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
// tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
// tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)

// fmt.Println(tree)

// schema := schema.RBTreeSchema(*tree)

// // data, err := os.ReadFile("in/rbt.json")
// // if err != nil {
// // 	log.Fatalf("Error reading file: %v", err)
// // 	return
// // }

// // bst, err := schema.Unmarshal(data)
// // if err != nil {
// // 	fmt.Println("Error unmarshalling BST:", err)
// // 	return
// // }
// dotGraph := schema.Dotify()
// fmt.Println(dotGraph.String())

// err := cmdMake(dotGraph, []string{"-png"})
// if err != nil {
// 	fmt.Println("Error:", err)
// 	return
// }
