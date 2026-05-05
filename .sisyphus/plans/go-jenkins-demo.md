# Go Minimal Web App for Jenkins Learning

## TL;DR

> **Quick Summary**: Tạo project Go nhỏ nhất có thể - 1 file main.go chạy web server trả về HTML đơn giản, kèm Dockerfile để deploy. không có Jenkinsfile - bạn tự viết sau.
> 
> **Deliverables**:
> - Go 1.26.1 installed và cấu hình
> - main.go - web server minimal hiển thị trang HTML
> - Dockerfile - đa giai đoạn build nhỏ gọn
> - Git repo initialized
> 
> **Estimated Effort**: Quick
> **Parallel Execution**: YES - 1 wave + verification
> **Critical Path**: Task 1 → Task 2 → Task 3

---

## Context

### Original Request
User cần 1 project Go nhỏ gọn để học Jenkins. Chỉ cần FE đơn giản hoặc cái gì đó biết nó chạy. Tự viết Jenkinsfile sau.

### Interview Summary
**Key Discussions**:
- Ban đầu: REST API + comprehensive Jenkins pipelines → User clarified: QUÁ NHIỀU, chỉ cần nhỏ nhất có thể
- User sẽ tự viết Jenkinsfile → KHÔNG tạo sẵn Jenkinsfile
- Cần visual confirmation (web page) để biết app đang chạy
- Đã có Jenkins, chỉ cần project

### Metis Review
**Identified Gaps** (addressed):
- Scope quá lớn ban đầu → Thu hẹp còn minimal
- Không cần framework (Gin) → Stdlib net/http đủ
- Không cần Docker registry → Chỉ Dockerfile local build

---

## Work Objectives

### Core Objective
Tạo project Go minimal nhất để làm substrate cho Jenkins pipeline practice.

### Concrete Deliverables
- Go 1.26.1 installed tại /usr/local/go
- `/home/dat/Code/prj_jenkin/golang/main.go` - HTTP server trả HTML page
- `/home/dat/Code/prj_jenkin/golang/Dockerfile` - multi-stage build
- `/home/dat/Code/prj_jenkin/golang/go.mod` - Go module init
- Git repo initialized với initial commit

### Definition of Done
- [ ] `go version` trả về go1.26.1
- [ ] `go run main.go` start server, `curl localhost:8080` trả HTML page
- [ ] `docker build -t go-demo .` success
- [ ] `docker run -p 8080:8080 go-demo` chạy và truy cập được
- [ ] `git log` có initial commit

### Must Have
- Go 1.26.1 hoạt động
- 1 web server hiển thị trang HTML khi truy cập
- Dockerfile hoạt động
- Git repo

### Must NOT Have (Guardrails)
- KHÔNG tạo Jenkinsfile (user tự viết)
- KHÔNG dùng framework bên ngoài (Gin, Echo, v.v.)
- KHÔNG tạo database hay storage
- KHÔNG tạo cấu trúc thư mục phức tạp (cmd/, internal/, pkg/)
- KHÔNG over-engineer - giữ tối thiểu

---

## Verification Strategy (MANDATORY)

> **ZERO HUMAN INTERVENTION** - ALL verification is agent-executed.

### Test Decision
- **Infrastructure exists**: NO
- **Automated tests**: None (project quá nhỏ, tự test bằng curl)
- **Framework**: none

### QA Policy
- **Verification**: Bash (curl) - Start server, send request, assert response

---

## Execution Strategy

### Parallel Execution Waves

```
Wave 1 (Foundation - sequential, each depends on previous):
├── Task 1: Install Go 1.26.1 + configure PATH [quick]
├── Task 2: Create main.go + go.mod [quick]
└── Task 3: Create Dockerfile + .gitignore [quick]

Wave FINAL (Verification):
├── Task F1: Plan compliance audit (oracle)
├── Task F2: Code quality review (unspecified-high)
├── Task F3: Real manual QA (unspecified-high)
└── Task F4: Scope fidelity check (deep)
```

### Dependency Matrix

- **1**: None → 2, 3
- **2**: 1 → 3
- **3**: 2 → F1-F4

### Agent Dispatch Summary

- **Wave 1**: Task 1 → `quick`, Task 2 → `quick`, Task 3 → `quick`
- **FINAL**: F1 → `oracle`, F2 → `unspecified-high`, F3 → `unspecified-high`, F4 → `deep`

---

## TODOs

- [ ] 1. Install Go 1.26.1 + Configure PATH

  **What to do**:
  - Download go1.26.1.linux-amd64.tar.gz from https://go.dev/dl/
  - Remove any existing Go installation at /usr/local/go
  - Extract tarball to /usr/local/go
  - Add `/usr/local/go/bin` and `$HOME/go/bin` to PATH in ~/.zshrc (or ~/.bashrc)
  - Source the shell config
  - Verify: `go version` returns go1.26.1

  **Must NOT do**:
  - Do NOT install via apt (outdated version)
  - Do NOT modify system-wide config beyond /usr/local/go
  - Do NOT install additional Go tools yet

  **Recommended Agent Profile**:
  - **Category**: `quick`
    - Reason: Standard installation task, well-documented steps
  - **Skills**: []

  **Parallelization**:
  - **Can Run In Parallel**: NO
  - **Parallel Group**: Sequential
  - **Blocks**: Task 2, Task 3
  - **Blocked By**: None

  **References**:
  - Official Go install: https://go.dev/doc/install
  - Shell config: `~/.zshrc` or `~/.bashrc` (zsh is default on this system)

  **Acceptance Criteria**:
  - [ ] `go version` outputs `go1.26.1 linux/amd64`
  - [ ] `which go` returns `/usr/local/go/bin/go`
  - [ ] `go env GOPATH` returns `/home/dat/go`

  **QA Scenarios (MANDATORY)**:

  ```
  Scenario: Go installation works
    Tool: Bash
    Preconditions: Go installed and PATH configured
    Steps:
      1. Run: go version
      2. Assert output contains "go1.26.1"
      3. Run: which go
      4. Assert output is "/usr/local/go/bin/go"
    Expected Result: Go 1.26.1 accessible from terminal
    Evidence: .sisyphus/evidence/task-1-go-version.txt

  Scenario: PATH persists in new shell
    Tool: Bash
    Preconditions: Shell config updated
    Steps:
      1. Run: zsh -c 'go version'
      2. Assert output contains "go1.26.1"
    Expected Result: Go works in fresh shell session
    Evidence: .sisyphus/evidence/task-1-path-persists.txt
  ```

  **Commit**: NO (not project files)

---

- [ ] 2. Create main.go + go.mod

  **What to do**:
  - Run `go mod init go-demo` in /home/dat/Code/prj_jenkin/golang
  - Create main.go with:
    - `package main`
    - import "net/http" and "fmt" only (no external deps)
    - Handler cho "/" trả về HTML page đơn giản với:
      - Tiêu đề: "Go Demo - Jenkins Learning"
      - Hiển thị hostname (os.Hostname())
      - Hiển thị thời gian server start
      - Hiển thị version (hardcode "1.0.0")
      - CSS inline tối giản (dark theme, centered)
    - Handler cho "/health" trả về JSON `{"status":"ok"}`
    - Listen trên ":8080"
  - Run `go build -o go-demo .` để verify syntax đúng
  - Run server briefly và test `curl localhost:8080` + `curl localhost:8080/health`

  **Must NOT do**:
  - Do NOT import any external packages (gin, echo, chi, etc.)
  - Do NOT create subdirectories (cmd/, internal/, pkg/)
  - Do NOT add unit tests (project quá nhỏ)
  - Do NOT create config files beyond go.mod

  **Recommended Agent Profile**:
  - **Category**: `quick`
    - Reason: Single file creation, straightforward
  - **Skills**: []

  **Parallelization**:
  - **Can Run In Parallel**: NO (depends on Go being installed)
  - **Parallel Group**: Sequential (after Task 1)
  - **Blocks**: Task 3
  - **Blocked By**: Task 1

  **References**:
  - Working directory: `/home/dat/Code/prj_jenkin/golang`
  - Go stdlib: https://pkg.go.dev/net/http
  - Go stdlib: https://pkg.go.dev/os#Hostname

  **Acceptance Criteria**:
  - [ ] `go build` succeeds without errors
  - [ ] `curl localhost:8080` returns HTML page
  - [ ] `curl localhost:8080/health` returns `{"status":"ok"}`
  - [ ] HTML page contains "Go Demo" and hostname
  - [ ] No external dependencies in go.mod (only stdlib)

  **QA Scenarios (MANDATORY)**:

  ```
  Scenario: Web server serves HTML page
    Tool: Bash
    Preconditions: Server running on port 8080
    Steps:
      1. Start server: go run main.go &
      2. Run: curl -s localhost:8080
      3. Assert output contains "Go Demo"
      4. Assert output contains "Jenkins Learning"
      5. Kill server process
    Expected Result: HTML page with app info rendered correctly
    Evidence: .sisyphus/evidence/task-2-html-response.txt

  Scenario: Health endpoint returns JSON
    Tool: Bash
    Preconditions: Server running on port 8080
    Steps:
      1. Start server: go run main.go &
      2. Run: curl -s localhost:8080/health
      3. Assert output contains "ok"
      4. Assert Content-Type contains "application/json"
      5. Kill server process
    Expected Result: JSON health response
    Evidence: .sisyphus/evidence/task-2-health-response.txt
  ```

  **Commit**: YES
  - Message: `feat: init minimal go web server`
  - Files: `main.go`, `go.mod`, `go.sum`
  - Pre-commit: `go build`

---

- [ ] 3. Create Dockerfile + .gitignore + Git Init

  **What to do**:
  - Create Dockerfile (multi-stage build):
    ```dockerfile
    # Build stage
    FROM golang:1.26-alpine AS builder
    WORKDIR /app
    COPY go.mod go.sum ./
    RUN go mod download
    COPY . .
    RUN CGO_ENABLED=0 GOOS=linux go build -o go-demo

    # Run stage
    FROM alpine:latest
    RUN apk --no-cache add ca-certificates
    WORKDIR /root/
    COPY --from=builder /app/go-demo .
    EXPOSE 8080
    CMD ["./go-demo"]
    ```
  - Create .gitignore:
    ```
    go-demo
    *.exe
    ```
  - Run `git init`
  - Run `git add .`
  - Run `git commit -m "feat: init go-demo project with dockerfile"`
  - Run `docker build -t go-demo .` để verify
  - (Optional) Run `docker run -d -p 8080:8080 go-demo` và test curl

  **Must NOT do**:
  - Do NOT create Jenkinsfile (user sẽ tự viết)
  - Do NOT push to remote repo (user tự setup)
  - Do NOT create docker-compose.yml (over-engineering)
  - Do NOT add Makefile (keep minimal)

  **Recommended Agent Profile**:
  - **Category**: `quick`
    - Reason: Simple Dockerfile + git init
  - **Skills**: []

  **Parallelization**:
  - **Can Run In Parallel**: NO (depends on main.go existing)
  - **Parallel Group**: Sequential (after Task 2)
  - **Blocks**: F1-F4
  - **Blocked By**: Task 2

  **References**:
  - Working directory: `/home/dat/Code/prj_jenkin/golang`
  - Docker multi-stage: https://docs.docker.com/build/building/multi-stage/

  **Acceptance Criteria**:
  - [ ] `docker build -t go-demo .` succeeds
  - [ ] `docker run --rm -p 8080:8080 go-demo` starts and serves HTML
  - [ ] `curl localhost:8080` works inside container
  - [ ] Git repo initialized with initial commit
  - [ ] .gitignore excludes binary

  **QA Scenarios (MANDATORY)**:

  ```
  Scenario: Docker build and run succeeds
    Tool: Bash
    Preconditions: Docker available
    Steps:
      1. Run: docker build -t go-demo .
      2. Assert build succeeds (exit code 0)
      3. Run: docker run -d -p 8080:8080 --name go-demo-test go-demo
      4. Sleep 2 seconds
      5. Run: curl -s localhost:8080
      6. Assert output contains "Go Demo"
      7. Run: docker stop go-demo-test && docker rm go-demo-test
    Expected Result: Containerized app serves HTML page
    Failure Indicators: Build fails, container exits immediately, curl returns nothing
    Evidence: .sisyphus/evidence/task-3-docker-run.txt

  Scenario: Git repo has initial commit
    Tool: Bash
    Preconditions: Git init done
    Steps:
      1. Run: git log --oneline
      2. Assert output contains "feat: init"
      3. Run: git status
      4. Assert "nothing to commit" or "working tree clean"
    Expected Result: Clean git repo with initial commit
    Evidence: .sisyphus/evidence/task-3-git-status.txt
  ```

  **Commit**: YES (grouped with Task 2's commit)
  - Message: `feat: add dockerfile and git init`
  - Files: `Dockerfile`, `.gitignore`
  - Pre-commit: `docker build -t go-demo .`

---

## Final Verification Wave (MANDATORY — after ALL implementation tasks)

- [ ] F1. **Plan Compliance Audit** — `oracle`
  Read the plan end-to-end. Verify: Go installed, main.go serves HTML + /health, Dockerfile builds, git initialized. Check evidence in .sisyphus/evidence/. Compare deliverables against plan.
  Output: `Must Have [4/4] | Must NOT Have [0 violations] | Tasks [3/3] | VERDICT: APPROVE/REJECT`

- [ ] F2. **Code Quality Review** — `unspecified-high`
  Run `go vet ./...` and `go build`. Review main.go for: proper error handling, no unused imports, clean code. Check Dockerfile for best practices. Check .gitignore completeness.
  Output: `Build [PASS/FAIL] | Vet [PASS/FAIL] | Files [3 clean/0 issues] | VERDICT`

- [ ] F3. **Real Manual QA** — `unspecified-high`
  Start server: `go run main.go`. Test `curl localhost:8080` HTML response. Test `curl localhost:8080/health` JSON response. Docker build and run. Verify container serves same content. Kill all processes. Save evidence to `.sisyphus/evidence/final-qa/`.
  Output: `Scenarios [4/4 pass] | Integration [2/2] | Edge Cases [1 tested] | VERDICT`

- [ ] F4. **Scope Fidelity Check** — `deep`
  Verify: Only 4 files exist (main.go, go.mod, Dockerfile, .gitignore). No Jenkinsfile exists. No external dependencies. No subdirectories. No over-engineering.
  Output: `Tasks [3/3 compliant] | Contamination [CLEAN] | Unaccounted [CLEAN] | VERDICT`

---

## Commit Strategy

- **Task 2+3 (combined)**: `feat: init go-demo project with dockerfile` - main.go, go.mod, go.sum, Dockerfile, .gitignore

---

## Success Criteria

### Verification Commands
```bash
go version                                    # Expected: go1.26.1
go build -o go-demo .                        # Expected: success
curl -s localhost:8080 | grep "Go Demo"      # Expected: match found
curl -s localhost:8080/health                 # Expected: {"status":"ok"}
docker build -t go-demo .                    # Expected: success
git log --oneline                            # Expected: initial commit
```

### Final Checklist
- [ ] Go 1.26.1 installed and working
- [ ] Web server serves HTML page on :8080
- [ ] /health endpoint returns JSON
- [ ] Docker build succeeds
- [ ] Docker container runs and serves content
- [ ] Git repo initialized with commit
- [ ] NO Jenkinsfile exists (explicit exclusion)
- [ ] NO external dependencies (stdlib only)