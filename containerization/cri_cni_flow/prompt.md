**Main Goal**
- Trigger a Mock CNI Plugin by running a containerd container.

**Context**
- The main goal of this task should be represented as a Ginkgo Golang test.
- Create a very simple CNI Mock that only logs calls from the outside world, i.e., containerd.
- What the test should look like:
    1. Configure containerd for the test using configuration files.
    2. Run a container using the containerd SDK.
    3. This execution should trigger the CNI Mock.
- All work should be placed in the `./containerization/cri_cni` directory.

**Workflow Rules**
- Do not provide the entire solution at once.
- Do not modify any files or run commands; only propose them to me.
- You may ask clarifying questions.
- You may propose better solutions than those in the Context.
- Provide a high-level plan before proposing any concrete coding steps.
- Once I confirm the high-level plan, break it down into small steps.
- Propose each small step one by one, and do not proceed to the next step until I confirm the current one.
- When proposing a small step, provide the piece of code that implements that step.

**Output Rules**
- You may use emojis.
- You may use numbered and bulleted lists.
- You may use headers and separators.
- You may use diagrams.