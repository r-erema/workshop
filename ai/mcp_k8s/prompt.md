**Main Goal**
- Understand how an LLM interacts with an MCP server.

**Context**
- The main goal of this task should be represented as a Ginkgo Golang test.
- We already have an MCP K8s server; you should use and examine it at `/home/erema/h/workshop/ai/mcp_k8s/pkg/mcp/server.go`.
- We should use an MCP client from the package `github.com/mark3labs/mcp-go`, if possible.
- For the K8s test API server, we should use the package `sigs.k8s.io/controller-runtime/pkg/envtest`.
- For the LLM, we should use Ollama, run via the `github.com/testcontainers/testcontainers-go` library, and select the smallest available LLM model in Ollama for test purposes.
- What the test should look like:
    1. Run the K8s API server using `sigs.k8s.io/controller-runtime/pkg/envtest`.
    2. Run a Testcontainer with Ollama.
    3. The test will include a prompt such as "list all pods in default namespace", which will be passed to the LLM.
    4. The actors (K8s API Server, MCP Client, MCP Server, LLM) interact with each other, ultimately list K8S pods.
    5. Test should in the file `/home/erema/h/workshop/ai/mcp_k8s/e2e_test.go`

**Workflow Rules**
- Do not provide the entire solution at once.
- You may ask clarifying questions.
- You may propose better solutions than those in the Context.
- Provide a high-level plan before proposing any concrete coding steps.
- Once I confirm the high-level plan, break it down into small steps.
- Propose each small step one by one, and do not proceed to the next step until I confirm the current one.
- When proposing a small step, provide the piece of code implementing that step.

**Output Rules**
- You may use emojis.
- You may use numbered and bulleted lists.
- You may use headers and separators.
- You may use diagrams.