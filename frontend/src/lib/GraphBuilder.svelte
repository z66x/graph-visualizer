<script lang="ts">
  let {
    nodes,
    directed,
    onaddnode,
    onaddedge,
    onremovenode,
    onclear,
    ontoggedirected
  }: {
    nodes: string[];
    directed: boolean;
    onaddnode: (name: string) => void;
    onaddedge: (from: string, to: string, weight: number) => void;
    onremovenode: (name: string) => void;
    onclear: () => void;
    ontoggedirected: () => void;
  } = $props();

  let newNode = $state('');
  let edgeFrom = $state('');
  let edgeTo = $state('');
  let edgeWeight = $state(1);

  function handleAddNode() {
    onaddnode(newNode.trim().toUpperCase());
    newNode = '';
  }

  function handleAddEdge() {
    onaddedge(edgeFrom, edgeTo, edgeWeight);
  }
</script>

<div class="builder">
  <div class="row">
    <div class="group">
      <label for="newnode">Add Node</label>
      <div class="inline">
        <input
          id="newnode"
          bind:value={newNode}
          placeholder="e.g. A"
          maxlength="3"
          onkeydown={(e) => e.key === 'Enter' && handleAddNode()}
        />
        <button onclick={handleAddNode}>Add</button>
      </div>
    </div>

    <div class="group">
      <span class="groupLabel">Mode</span>
      <button class="toggle" onclick={ontoggedirected}>
        {directed ? 'Directed' : 'Undirected'}
      </button>
    </div>

    <div class="group">
      <span class="groupLabel">Graph</span>
      <button class="danger" onclick={onclear}>Clear All</button>
    </div>
  </div>

<div class="group">
  <label for="edgefrom">Add Edge</label>
  <div class="inline">
    <select id="edgefrom" bind:value={edgeFrom}>
      <option value="">From</option>
      {#each nodes as n}
        <option value={n}>{n}</option>
      {/each}
    </select>
    <select id="edgeto" bind:value={edgeTo}>
      <option value="">To</option>
      {#each nodes as n}
        <option value={n}>{n}</option>
      {/each}
    </select>
    <input id="edgeweight" type="number" bind:value={edgeWeight} min="1" max="99" style="width:60px" />
    <button onclick={handleAddEdge}>Add</button>
  </div>
</div>

  {#if nodes.length > 0}
    <div class="nodelist">
      {#each nodes as n}
        <span class="nodetag">
          {n}
          <button class="x" onclick={() => onremovenode(n)}>×</button>
        </span>
      {/each}
    </div>
  {/if}
</div>

<style>
  .builder {
    padding: 1rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }
  .row {
    display: flex;
    gap: 1rem;
    flex-wrap: wrap;
  }
  .group {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }
  label, .groupLabel {
    font-size: 0.8rem;
    color: var(--primary-color);
    text-transform: uppercase;
  }
  .inline {
    display: flex;
    gap: 0.5rem;
    align-items: center;
  }
  input, select {
    background: transparent;
    color: var(--primary-color);
    border: 1px solid var(--secondary-color);
    border-radius: 6px;
    padding: 0.5rem 0.5rem;
    font-size: 1rem;
  }
  button {
    background: var(--accent-color);
    color: var(--primary-color);
    border: none;
    border-radius: 6px;
    padding: 0.5rem 1rem;
    cursor: pointer;
    font-size: 1rem;
  }
  button:hover { outline: 1px solid var(--primary-color); }
  .danger { background: #DD1100; }
  .nodelist {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
  }
  .nodetag {
    background-color: var(--secondary-color);
    border-radius: 16px;
    padding: 0.2rem 0.6rem;
    font-size: 1rem;
  }
  .x {
    background: none;
    color: var(--primary-color);
    padding: 0;
    font-size: 1rem;
  }
  .x:hover { color: #DD1100; outline: none;}
</style>