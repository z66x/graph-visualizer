<script lang="ts">
  import { GraphCanvas, Controls, GraphBuilder }  from '$lib';

  type Step = {
    visited: string[];
    current: string;
    frontier: string[];
    edgeTaken: Record<string, string>;
    message: string;
  };

  let nodes: string[] = [];
  let edges: Record<string, { to: string; weight: number }[]> = {};
  let directed = false;

  let steps: Step[] = [];
  let currentStep = 0;
  let playing = false;
  let speed = 800;
  let selectedAlgo = 'bfs';
  let startNode = '';
  let message = '';
  let interval: ReturnType<typeof setInterval>;

  function addNode(name: string) {
    if (!name || nodes.includes(name)) return;
    nodes = [...nodes, name];
    edges = { ...edges, [name]: [] };
    if (!startNode) startNode = name;
  }

  function addEdge(from: string, to: string, weight: number) {
    if (!from || !to || from === to) return;
    const already = edges[from]?.some(e => e.to === to);
    if (already) return;
    edges = {
      ...edges,
      [from]: [...(edges[from] ?? []), { to, weight }],
      ...(!directed && {
        [to]: [...(edges[to] ?? []), { to: from, weight }]
      })
    };
  }

  function removeNode(name: string) {
    nodes = nodes.filter(n => n !== name);
    const newEdges: typeof edges = {};
    for (const n of nodes) {
      newEdges[n] = (edges[n] ?? []).filter(e => e.to !== name);
    }
    edges = newEdges;
    if (startNode === name) startNode = nodes[0] ?? '';
    reset();
  }

  function clearGraph() {
    nodes = [];
    edges = {};
    startNode = '';
    reset();
  }

  async function runAlgorithm() {
    if (nodes.length === 0 || !startNode) return;
      const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/api/${selectedAlgo}`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        graph: { nodes, edges, directed },
        start: startNode
      })
    });
    const data = await res.json();
    steps = data.steps;
    currentStep = 0;
    message = steps[0]?.message ?? '';
  }

  function play() {
    if (steps.length === 0) return;
    playing = true;
    interval = setInterval(() => {
      if (currentStep < steps.length - 1) {
        currentStep++;
        message = steps[currentStep].message;
      } else {
        pause();
      }
    }, speed);
  }

  function pause() {
    playing = false;
    clearInterval(interval);
  }

  function reset() {
    pause();
    steps = [];
    currentStep = 0;
    message = '';
  }
  function prev() {
    if (currentStep > 0) {
        pause();
        currentStep--;
        message = steps[currentStep].message;
    }
    }

    function next() {
    if (currentStep < steps.length - 1) {
        pause();
        currentStep++;
        message = steps[currentStep].message;
    }
}
  $: current = steps[currentStep] ?? null;
</script>

<main>
  <div class="layout">
  <div class="canvas">
  <GraphCanvas
        {nodes}
        {edges}
        {directed}
        visited={current?.visited ?? []}
        frontier={current?.frontier ?? []}
        currentNode={current?.current ?? ''}
        onremovenode={removeNode}
        stepInfo={steps.length > 0 ? `Step ${currentStep + 1} / ${steps.length}` : ''}
      />
    </div>
    <div class="right">
    <h1>Graph Algorithm Visualizer</h1>
      <GraphBuilder
        {nodes}
        {directed}
        onaddnode={addNode}
        onaddedge={addEdge}
        onremovenode={removeNode}
        onclear={clearGraph}
        ontoggedirected={() => { directed = !directed; reset(); }}
      />
    <Controls
    bind:selectedAlgo
    bind:startNode
    bind:speed
    {nodes}
    {playing}
    {message}
    {currentStep}
    totalSteps={steps.length}
    onrun={runAlgorithm}
    onplay={play}
    onpause={pause}
    onprev={prev}
    onnext={next}
    onreset={reset}
    />
    </div>
  </div>
</main>

<style>
    main {
        padding: 2rem;
        height: 100vh;
    }
  h1 {
    text-align: center;
    color: var(--accent-color);
    font-size: 2rem;
    padding-bottom: 1rem;
    border-bottom: 1px solid var(--secondary-color);
  }
  .layout {
    display: flex;
    height: 100%;
    gap: 2rem;
    justify-content: space-between;
  }
  .canvas {
    flex: 1;
    border: 2px solid var(--secondary-color);
  }
  .right {
    display: flex;
    width: min(400px, 40%);
    flex-direction: column;
    gap: 1rem;
  }
</style>