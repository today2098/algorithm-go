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
  attributes:
    PROBLEM: https://onlinejudge.u-aizu.ac.jp/courses/library/3/DSL/1/DSL_1_A
  bundledCode: "Traceback (most recent call last):\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/documentation/build.py\"\
    , line 71, in _render_source_code_stat\n    bundled_code = language.bundle(stat.path,\
    \ basedir=basedir, options={'include_paths': [basedir]}).decode()\n  File \"/home/runner/.local/lib/python3.10/site-packages/onlinejudge_verify/languages/user_defined.py\"\
    , line 68, in bundle\n    raise RuntimeError('bundler is not specified: {}'.format(str(path)))\n\
    RuntimeError: bundler is not specified: test/aoj-DSL_1_A-union_find.test.go\n"
  code: "//go:build ignore\n\n// verification-helper: PROBLEM https://onlinejudge.u-aizu.ac.jp/courses/library/3/DSL/1/DSL_1_A\n\
    \npackage main\n\nimport (\n\t\"bufio\"\n\t\"fmt\"\n\t\"math\"\n\t\"os\"\n\t\"\
    strconv\"\n\n\t\"github.com/today2098/algorithm-go/algorithm\"\n)\n\nvar sc =\
    \ bufio.NewScanner(os.Stdin)\nvar wr = bufio.NewWriter(os.Stdout)\n\nfunc getInt()\
    \ int {\n\tsc.Scan()\n\telem, err := strconv.Atoi(sc.Text())\n\tif err != nil\
    \ {\n\t\tpanic(err)\n\t}\n\treturn elem\n}\n\nfunc out(x ...any) {\n\tfmt.Fprintln(wr,\
    \ x...)\n}\n\nfunc main() {\n\tsc.Split(bufio.ScanWords)\n\tsc.Buffer([]byte{},\
    \ math.MaxInt32)\n\tdefer wr.Flush()\n\n\tn, q := getInt(), getInt()\n\n\tuf :=\
    \ algorithm.NewUnionFind(n)\n\tfor i := 0; i < q; i++ {\n\t\tcom, x, y := getInt(),\
    \ getInt(), getInt()\n\n\t\tif com == 0 {\n\t\t\tuf.Unite(x, y)\n\t\t} else {\n\
    \t\t\tif uf.IsSame(x, y) {\n\t\t\t\tout(1)\n\t\t\t} else {\n\t\t\t\tout(0)\n\t\
    \t\t}\n\t\t}\n\t}\n}\n"
  dependsOn:
  - main.go
  - test/aoj-GRL_1_A-dijkstra.test.go
  - test/aoj-ALDS1_3_C-deque.test.go
  - test/aoj-ALDS1_3_B-queue.test.go
  - test/aoj-ALDS1_9_C-binary_haep.test.go
  - test/aoj-GRL_1_B-bellman_ford.test.go
  - test/aoj-ITP1_1_A.test.go
  - test/aoj-ALDS1_3_A-stack.test.go
  - algorithm/bellman_ford.go
  - algorithm/deque.go
  - algorithm/union_find.go
  - algorithm/dijkstra.go
  - algorithm/queue.go
  - algorithm/binary_heap.go
  - algorithm/stack.go
  isVerificationFile: true
  path: test/aoj-DSL_1_A-union_find.test.go
  requiredBy:
  - main.go
  - algorithm/bellman_ford.go
  - algorithm/deque.go
  - algorithm/union_find.go
  - algorithm/dijkstra.go
  - algorithm/queue.go
  - algorithm/binary_heap.go
  - algorithm/stack.go
  timestamp: '2024-08-25 05:23:23+09:00'
  verificationStatus: TEST_ACCEPTED
  verifiedWith:
  - test/aoj-GRL_1_A-dijkstra.test.go
  - test/aoj-ALDS1_3_C-deque.test.go
  - test/aoj-ALDS1_3_B-queue.test.go
  - test/aoj-ALDS1_9_C-binary_haep.test.go
  - test/aoj-GRL_1_B-bellman_ford.test.go
  - test/aoj-ITP1_1_A.test.go
  - test/aoj-ALDS1_3_A-stack.test.go
documentation_of: test/aoj-DSL_1_A-union_find.test.go
layout: document
redirect_from:
- /verify/test/aoj-DSL_1_A-union_find.test.go
- /verify/test/aoj-DSL_1_A-union_find.test.go.html
title: test/aoj-DSL_1_A-union_find.test.go
---
