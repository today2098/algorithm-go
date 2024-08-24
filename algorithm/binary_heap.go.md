---
data:
  _extendedDependsOn:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque.go
    title: algorithm/deque.go
  - icon: ':heavy_check_mark:'
    path: algorithm/dijkstra.go
    title: algorithm/dijkstra.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue.go
    title: algorithm/queue.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack.go
    title: algorithm/stack.go
  - icon: ':heavy_check_mark:'
    path: algorithm/union_find.go
    title: algorithm/union_find.go
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ALDS1_3_A-stack.test.go
    title: test/aoj-ALDS1_3_A-stack.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ALDS1_3_B-queue.test.go
    title: test/aoj-ALDS1_3_B-queue.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ALDS1_3_C-deque.test.go
    title: test/aoj-ALDS1_3_C-deque.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ALDS1_9_C-binary_haep.test.go
    title: test/aoj-ALDS1_9_C-binary_haep.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-DSL_1_A-union_find.test.go
    title: test/aoj-DSL_1_A-union_find.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-GRL_1_A-dijkstra.test.go
    title: test/aoj-GRL_1_A-dijkstra.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _extendedRequiredBy:
  - icon: ':heavy_check_mark:'
    path: algorithm/deque.go
    title: algorithm/deque.go
  - icon: ':heavy_check_mark:'
    path: algorithm/dijkstra.go
    title: algorithm/dijkstra.go
  - icon: ':heavy_check_mark:'
    path: algorithm/queue.go
    title: algorithm/queue.go
  - icon: ':heavy_check_mark:'
    path: algorithm/stack.go
    title: algorithm/stack.go
  - icon: ':heavy_check_mark:'
    path: algorithm/union_find.go
    title: algorithm/union_find.go
  - icon: ':heavy_check_mark:'
    path: main.go
    title: main.go
  _extendedVerifiedWith:
  - icon: ':heavy_check_mark:'
    path: test/aoj-ALDS1_3_A-stack.test.go
    title: test/aoj-ALDS1_3_A-stack.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ALDS1_3_B-queue.test.go
    title: test/aoj-ALDS1_3_B-queue.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ALDS1_3_C-deque.test.go
    title: test/aoj-ALDS1_3_C-deque.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ALDS1_9_C-binary_haep.test.go
    title: test/aoj-ALDS1_9_C-binary_haep.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-DSL_1_A-union_find.test.go
    title: test/aoj-DSL_1_A-union_find.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-GRL_1_A-dijkstra.test.go
    title: test/aoj-GRL_1_A-dijkstra.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _isVerificationFailed: false
  _pathExtension: go
  _verificationStatusIcon: ':heavy_check_mark:'
  attributes: {}
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: algorithm/binary_heap.go\n"
  code: "package algorithm\n\nimport (\n\t\"errors\"\n\n\t\"golang.org/x/exp/constraints\"\
    \n)\n\nvar ErrBinaryHeapEmpty = errors.New(\"BinaryHeap: binary-heap is empty\"\
    )\n\ntype BinaryHeap[T any] struct {\n\tcomp BinaryHeapCompFunc[T]\n\ttree []T\n\
    }\n\ntype BinaryHeapCompFunc[T any] func(a, b T) bool\n\nfunc NewBinaryHeap[T\
    \ any](f BinaryHeapCompFunc[T]) *BinaryHeap[T] {\n\treturn &BinaryHeap[T]{\n\t\
    \tcomp: f,\n\t\ttree: make([]T, 1),\n\t}\n}\n\nfunc NewDefaultBinaryHeap[T constraints.Ordered]()\
    \ *BinaryHeap[T] {\n\treturn NewBinaryHeap(func(a, b T) bool {\n\t\treturn a >\
    \ b\n\t})\n}\n\nfunc (b *BinaryHeap[T]) shiftUp(i int) {\n\tp := i >> 1\n\tfor\
    \ 1 <= p {\n\t\tif b.comp(b.tree[p], b.tree[i]) {\n\t\t\tbreak\n\t\t}\n\t\tb.tree[p],\
    \ b.tree[i] = b.tree[i], b.tree[p]\n\t\ti = p\n\t\tp >>= 1\n\t}\n}\n\nfunc (b\
    \ *BinaryHeap[T]) shiftDown(i int) {\n\tl, r := i<<1, i<<1|1\n\tfor l <= b.Size()\
    \ {\n\t\tif b.Size() < r || b.comp(b.tree[l], b.tree[r]) {\n\t\t\tif b.comp(b.tree[i],\
    \ b.tree[l]) {\n\t\t\t\tbreak\n\t\t\t}\n\t\t\tb.tree[i], b.tree[l] = b.tree[l],\
    \ b.tree[i]\n\t\t\ti = l\n\t\t} else {\n\t\t\tif b.comp(b.tree[i], b.tree[r])\
    \ {\n\t\t\t\tbreak\n\t\t\t}\n\t\t\tb.tree[i], b.tree[r] = b.tree[r], b.tree[i]\n\
    \t\t\ti = r\n\t\t}\n\t\tl, r = i<<1, i<<1|1\n\t}\n}\n\nfunc (b *BinaryHeap[T])\
    \ Empty() bool {\n\treturn b.Size() == 0\n}\n\nfunc (b *BinaryHeap[T]) Size()\
    \ int {\n\treturn len(b.tree) - 1\n}\n\nfunc (b *BinaryHeap[T]) Top() T {\n\t\
    if b.Empty() {\n\t\tpanic(ErrBinaryHeapEmpty)\n\t}\n\treturn b.tree[1]\n}\n\n\
    func (b *BinaryHeap[T]) Push(x T) {\n\tb.tree = append(b.tree, x)\n\tb.shiftUp(b.Size())\n\
    }\n\nfunc (b *BinaryHeap[T]) Pop() T {\n\tif b.Empty() {\n\t\tpanic(ErrBinaryHeapEmpty)\n\
    \t}\n\tres := b.tree[1]\n\tb.tree[1] = b.tree[b.Size()]\n\tb.tree = b.tree[:len(b.tree)-1]\n\
    \tif !b.Empty() {\n\t\tb.shiftDown(1)\n\t}\n\treturn res\n}\n"
  dependsOn:
  - main.go
  - test/aoj-GRL_1_A-dijkstra.test.go
  - test/aoj-ALDS1_3_C-deque.test.go
  - test/aoj-ALDS1_3_B-queue.test.go
  - test/aoj-ALDS1_9_C-binary_haep.test.go
  - test/aoj-ITP1_1_A.test.go
  - test/aoj-DSL_1_A-union_find.test.go
  - test/aoj-ALDS1_3_A-stack.test.go
  - algorithm/deque.go
  - algorithm/union_find.go
  - algorithm/dijkstra.go
  - algorithm/queue.go
  - algorithm/stack.go
  isVerificationFile: false
  path: algorithm/binary_heap.go
  requiredBy:
  - main.go
  - algorithm/deque.go
  - algorithm/union_find.go
  - algorithm/dijkstra.go
  - algorithm/queue.go
  - algorithm/stack.go
  timestamp: '2024-08-25 00:57:10+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - test/aoj-GRL_1_A-dijkstra.test.go
  - test/aoj-ALDS1_3_C-deque.test.go
  - test/aoj-ALDS1_3_B-queue.test.go
  - test/aoj-ALDS1_9_C-binary_haep.test.go
  - test/aoj-ITP1_1_A.test.go
  - test/aoj-DSL_1_A-union_find.test.go
  - test/aoj-ALDS1_3_A-stack.test.go
documentation_of: algorithm/binary_heap.go
layout: document
redirect_from:
- /library/algorithm/binary_heap.go
- /library/algorithm/binary_heap.go.html
title: algorithm/binary_heap.go
---
