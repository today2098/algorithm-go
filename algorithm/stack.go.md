---
data:
  _extendedDependsOn:
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
    RuntimeError: bundler is not specified: algorithm/stack.go\n"
  code: "package algorithm\n\nimport \"errors\"\n\nvar ErrStackEmpty = errors.New(\"\
    Stack: stack is empty\")\n\ntype Stack[T any] struct {\n\tData []T\n}\n\nfunc\
    \ NewStack[T any]() *Stack[T] {\n\treturn &Stack[T]{Data: []T{}}\n}\n\nfunc (s\
    \ *Stack[T]) Empty() bool {\n\treturn s.Size() == 0\n}\n\nfunc (s *Stack[T]) Size()\
    \ int {\n\treturn len(s.Data)\n}\n\nfunc (s *Stack[T]) Top() T {\n\tif s.Empty()\
    \ {\n\t\tpanic(ErrStackEmpty)\n\t}\n\treturn s.Data[len(s.Data)-1]\n}\n\nfunc\
    \ (s *Stack[T]) Push(x T) {\n\ts.Data = append(s.Data, x)\n}\n\nfunc (s *Stack[T])\
    \ PushRange(v []T) {\n\tfor i := 0; i < len(v); i++ {\n\t\ts.Data = append(s.Data,\
    \ v[i])\n\t}\n}\n\nfunc (s *Stack[T]) Pop() T {\n\tif s.Empty() {\n\t\tpanic(ErrStackEmpty)\n\
    \t}\n\tres := s.Data[len(s.Data)-1]\n\ts.Data = s.Data[:len(s.Data)-1]\n\treturn\
    \ res\n}\n"
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
  - algorithm/binary_heap.go
  isVerificationFile: false
  path: algorithm/stack.go
  requiredBy:
  - main.go
  - algorithm/deque.go
  - algorithm/union_find.go
  - algorithm/dijkstra.go
  - algorithm/queue.go
  - algorithm/binary_heap.go
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
documentation_of: algorithm/stack.go
layout: document
redirect_from:
- /library/algorithm/stack.go
- /library/algorithm/stack.go.html
title: algorithm/stack.go
---
