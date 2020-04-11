class Node:
    def __init__(self, val):
        self.data = val
        self.left = None
        self.right = None

    def insert_node(self, val):
        if self.data:
            if val < self.data:
                if self.left is None:
                    self.left = Node(val)
                else:
                    self.left.insert_node(val)
            elif val > self.data:
                if self.right is None:
                    self.right = Node(val)
                else:
                    self.right.insert_node(val)
        else:
            self.data = val

    def print_tree(self):
        if self.left:
            self.left.print_tree()
        print(self.data)

        if self.right:
            self.right.print_tree()

    def height(self, node):
        if node is None:
            return 0

        else:
            l_depth = self.height(node.left)
            r_depth = self.height(node.right)

            if l_depth > r_depth:
                return l_depth + 1
            else:
                return r_depth + 1

root = Node(12)
root.insert_node(1000)
root.insert_node(1)
root.insert_node(3)
root.insert_node(13)
root.print_tree()
print(root.height(root))