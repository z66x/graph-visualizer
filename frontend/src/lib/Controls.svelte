<script lang="ts">
  let {
    selectedAlgo = $bindable(),
    startNode = $bindable(),
    speed = $bindable(),
    nodes,
    playing,
    message,
    currentStep,
    totalSteps,
    onrun,
    onplay,
    onpause,
    onprev,
    onnext,
    onreset
  }: {
    selectedAlgo: string;
    startNode: string;
    speed: number;
    nodes: string[];
    playing: boolean;
    message: string;
    currentStep: number;
    totalSteps: number;
    onrun: () => void;
    onplay: () => void;
    onpause: () => void;
    onprev: () => void;
    onnext: () => void;
    onreset: () => void;
  } = $props();
</script>

<div class="controls">
  <div class="group">
    <label for="algo">Algorithm</label>
    <select id="algo" bind:value={selectedAlgo}>
      <option value="bfs">BFS — Breadth First Search</option>
      <option value="dfs">DFS — Depth First Search</option>
      <option value="dijkstra">Dijkstra — Shortest Path</option>
    </select>
  </div>

  <div class="group">
    <label for="start">Start Node</label>
    <select id="start" bind:value={startNode}>
      <option value="">Select node</option>
      {#each nodes as node}
        <option value={node}>{node}</option>
      {/each}
    </select>
  </div>

  <div class="group">
    <label for="speed">Speed — {speed}ms per step</label>
    <input id="speed" type="range" min="200" max="2000" step="100" bind:value={speed} />
  </div>

    <div class="buttons">
    <button onclick={onrun}>Run</button>
    <button onclick={onprev} disabled={currentStep === 0 || totalSteps === 0}>‹</button>
    {#if !playing}
        <button onclick={onplay} disabled={totalSteps === 0}>▶</button>
    {:else}
        <button onclick={onpause}>⏸</button>
    {/if}
    <button onclick={onnext} disabled={currentStep >= totalSteps - 1 || totalSteps === 0}>›</button>
    <button onclick={onreset}>↺</button>
    </div>

  {#if message}
    <div class="message">> {message}</div>
  {/if}
</div>

<style>
  .controls {
    padding: 1rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }
  .group {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }
  label {
    font-size: 0.8rem;
    color: var(--primary-color);
    text-transform: uppercase;
  }
  input, select {
    background: transparent;
    color: var(--primary-color);
    border: 1px solid var(--secondary-color);
    border-radius: 6px;
    padding: 0.5rem 0.5rem;
    font-size: 1rem;
  }
  .buttons {
    display: flex;
    gap: 0.5rem;
    flex-wrap: wrap;
  }
  button {
    background: var(--accent-color);
    color: var(--primary-color);
    border: none;
    border-radius: 6px;
    padding: 0.5rem 1rem;
    cursor: pointer;
    font-size: 1rem;
    flex: 1;
  }
  button:hover { outline: 1px solid var(--primary-color); }
  button:disabled { background: var(--secondary-color); cursor: not-allowed; }
  .message {
    padding: 0.5rem;
    border-radius: 10px;
    font-size: 1rem;
    color: var(--primary-color);
  }
</style>