### Heap Sort

The heapsort algorithm can be divided into two parts.

In the first step, a heap is built out of the data (see Binary heap § Building a heap). The heap is often placed in an array with the layout of a complete binary tree. The complete binary tree maps the binary tree structure into the array indices; each array index represents a node; the index of the node's parent, left child branch, or right child branch are simple expressions. For a zero-based array, the root node is stored at index 0; if i is the index of the current node, then 
In the second step, a sorted array is created by repeatedly removing the largest element from the heap (the root of the heap), and inserting it into the array. The heap is updated after each removal to maintain the heap property. Once all objects have been removed from the heap, the result is a sorted array.

<p>
	<img src="https://russianblogs.com/images/19/2ed3e65e9cc6a624accec7380645dbeb.png"/>
</p>