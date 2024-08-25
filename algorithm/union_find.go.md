---
data:
  _extendedDependsOn:
  - icon: ':heavy_check_mark:'
    path: algorithm/bellman_ford.go
    title: algorithm/bellman_ford.go
  - icon: ':heavy_check_mark:'
    path: algorithm/binary_heap.go
    title: algorithm/binary_heap.go
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
    path: test/aoj-GRL_1_B-bellman_ford.test.go
    title: test/aoj-GRL_1_B-bellman_ford.test.go
  - icon: ':heavy_check_mark:'
    path: test/aoj-ITP1_1_A.test.go
    title: test/aoj-ITP1_1_A.test.go
  _extendedRequiredBy:
  - icon: ':heavy_check_mark:'
    path: algorithm/bellman_ford.go
    title: algorithm/bellman_ford.go
  - icon: ':heavy_check_mark:'
    path: algorithm/binary_heap.go
    title: algorithm/binary_heap.go
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
    path: test/aoj-GRL_1_B-bellman_ford.test.go
    title: test/aoj-GRL_1_B-bellman_ford.test.go
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
    RuntimeError: bundler is not specified: algorithm/union_find.go\n"
  code: "package algorithm\n\nimport \"errors\"\n\nvar ErrUnionFindOutOfRange = errors.New(\"\
    UnionFind: index out of range\")\n\ntype UnionFind struct {\n\tpar []int\n\tgn\
    \  int\n}\n\nfunc NewUnionFind(n int) *UnionFind {\n\tpar := make([]int, n)\n\t\
    for i := 0; i < n; i++ {\n\t\tpar[i] = -1\n\t}\n\treturn &UnionFind{\n\t\tpar:\
    \ par,\n\t\tgn:  n,\n\t}\n}\n\nfunc (uf *UnionFind) Vn() int {\n\treturn len(uf.par)\n\
    }\n\nfunc (uf *UnionFind) Gn() int {\n\treturn uf.gn\n}\n\nfunc (uf *UnionFind)\
    \ Root(x int) int {\n\tif !(0 <= x && x < uf.Vn()) {\n\t\tpanic(ErrUnionFindOutOfRange)\n\
    \t}\n\tif uf.par[x] < 0 {\n\t\treturn x\n\t}\n\treturn uf.Root(uf.par[x])\n}\n\
    \nfunc (uf *UnionFind) Size(x int) int {\n\tif !(0 <= x && x < uf.Vn()) {\n\t\t\
    panic(ErrUnionFindOutOfRange)\n\t}\n\treturn -uf.par[uf.Root(x)]\n}\n\nfunc (uf\
    \ *UnionFind) IsSame(x, y int) bool {\n\tif !(0 <= x && x < uf.Vn() && 0 <= y\
    \ && y < uf.Vn()) {\n\t\tpanic(ErrUnionFindOutOfRange)\n\t}\n\treturn uf.Root(x)\
    \ == uf.Root(y)\n}\n\nfunc (uf *UnionFind) Unite(x, y int) bool {\n\tif !(0 <=\
    \ x && x < uf.Vn() && 0 <= y && y < uf.Vn()) {\n\t\tpanic(ErrUnionFindOutOfRange)\n\
    \t}\n\tx, y = uf.Root(x), uf.Root(y)\n\tif x == y {\n\t\treturn false\n\t}\n\t\
    if uf.Size(x) < uf.Size(y) {\n\t\tx, y = y, x\n\t}\n\tuf.par[x] += uf.par[y]\n\
    \tuf.par[y] = x\n\tuf.gn--\n\treturn true\n}\n\nfunc (uf *UnionFind) Reset() {\n\
    \tfor i := 0; i < uf.Vn(); i++ {\n\t\tuf.par[i] = -1\n\t}\n\tuf.gn = uf.Vn()\n\
    }\n"
  dependsOn:
  - main.go
  - test/aoj-GRL_1_A-dijkstra.test.go
  - test/aoj-ALDS1_3_C-deque.test.go
  - test/aoj-ALDS1_3_B-queue.test.go
  - test/aoj-ALDS1_9_C-binary_haep.test.go
  - test/aoj-GRL_1_B-bellman_ford.test.go
  - test/aoj-ITP1_1_A.test.go
  - test/aoj-DSL_1_A-union_find.test.go
  - test/aoj-ALDS1_3_A-stack.test.go
  - algorithm/bellman_ford.go
  - algorithm/deque.go
  - algorithm/dijkstra.go
  - algorithm/queue.go
  - algorithm/binary_heap.go
  - algorithm/stack.go
  isVerificationFile: false
  path: algorithm/union_find.go
  requiredBy:
  - main.go
  - algorithm/bellman_ford.go
  - algorithm/deque.go
  - algorithm/dijkstra.go
  - algorithm/queue.go
  - algorithm/binary_heap.go
  - algorithm/stack.go
  timestamp: '2024-08-25 05:23:23+09:00'
  verificationStatus: LIBRARY_ALL_AC
  verifiedWith:
  - test/aoj-GRL_1_A-dijkstra.test.go
  - test/aoj-ALDS1_3_C-deque.test.go
  - test/aoj-ALDS1_3_B-queue.test.go
  - test/aoj-ALDS1_9_C-binary_haep.test.go
  - test/aoj-GRL_1_B-bellman_ford.test.go
  - test/aoj-ITP1_1_A.test.go
  - test/aoj-DSL_1_A-union_find.test.go
  - test/aoj-ALDS1_3_A-stack.test.go
documentation_of: algorithm/union_find.go
layout: document
redirect_from:
- /library/algorithm/union_find.go
- /library/algorithm/union_find.go.html
title: algorithm/union_find.go
---
