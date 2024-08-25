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
    RuntimeError: bundler is not specified: algorithm/queue.go\n"
  code: "package algorithm\n\nimport \"errors\"\n\nvar ErrQueueEmpty = errors.New(\"\
    Queue: queue is empty\")\n\ntype Queue[T any] struct {\n\tData []T\n}\n\nfunc\
    \ NewQueue[T any]() *Queue[T] {\n\treturn &Queue[T]{Data: []T{}}\n}\n\nfunc (q\
    \ *Queue[T]) Empty() bool {\n\treturn q.Size() == 0\n}\n\nfunc (q *Queue[T]) Size()\
    \ int {\n\treturn len(q.Data)\n}\n\nfunc (q *Queue[T]) Front() T {\n\tif q.Empty()\
    \ {\n\t\tpanic(ErrQueueEmpty)\n\t}\n\treturn q.Data[0]\n}\n\nfunc (q *Queue[T])\
    \ Push(x T) {\n\tq.Data = append(q.Data, x)\n}\n\nfunc (q *Queue[T]) PushRange(v\
    \ []T) {\n\tfor i := 0; i < len(v); i++ {\n\t\tq.Data = append(q.Data, v[i])\n\
    \t}\n}\n\nfunc (q *Queue[T]) Pop() T {\n\tif q.Empty() {\n\t\tpanic(ErrQueueEmpty)\n\
    \t}\n\tres := q.Data[0]\n\tq.Data = q.Data[1:]\n\treturn res\n}\n"
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
  - algorithm/union_find.go
  - algorithm/dijkstra.go
  - algorithm/binary_heap.go
  - algorithm/stack.go
  isVerificationFile: false
  path: algorithm/queue.go
  requiredBy:
  - main.go
  - algorithm/bellman_ford.go
  - algorithm/deque.go
  - algorithm/union_find.go
  - algorithm/dijkstra.go
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
documentation_of: algorithm/queue.go
layout: document
redirect_from:
- /library/algorithm/queue.go
- /library/algorithm/queue.go.html
title: algorithm/queue.go
---
