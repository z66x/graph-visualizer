<script lang="ts">
  import * as d3 from 'd3';

  let {
    nodes,
    edges,
    visited,
    frontier,
    currentNode,
    directed = false,
    onremovenode = () => {},
    stepInfo = ''
  }: {
    nodes: string[];
    edges: Record<string, { to: string; weight: number }[]>;
    visited: string[];
    frontier: string[];
    currentNode: string;
    directed?: boolean;
    onremovenode?: (name: string) => void;
    stepInfo?: string;
  } = $props();

  let svgEl: SVGSVGElement | undefined = $state();
  let containerEl: HTMLDivElement | undefined = $state();
  let width = $state(560);
  let height = $state(500);
  const radius = 24;

  let positions: Record<string, { x: number; y: number }> = {};

  function initPosition(n: string, i: number, total: number) {
    if (positions[n]) return;
    const angle = (i / total) * 2 * Math.PI - Math.PI / 2;
    positions[n] = {
      x: width / 2 + 160 * Math.cos(angle),
      y: height / 2 + 160 * Math.sin(angle)
    };
  }

  function nodeColor(n: string) {
    if (n === currentNode) return '#f59e0b';
    if (frontier.includes(n)) return '#60a5fa';
    if (visited.includes(n)) return '#22c55e';
    return '#334155';
  }

  function draw() {
    if (!svgEl) return;

    // measure container and resize
    if (containerEl) {
      width = containerEl.clientWidth || 560;
      height = containerEl.clientHeight || 500;
    }

    const svg = d3.select(svgEl);
    svg.selectAll('*').remove();

    nodes.forEach((n, i) => initPosition(n, i, nodes.length));

    for (const key of Object.keys(positions)) {
      if (!nodes.includes(key)) delete positions[key];
    }

    for (const [from, targets] of Object.entries(edges)) {
      for (const { to, weight } of targets) {
        if (!positions[from] || !positions[to]) continue;

        const x1 = positions[from].x;
        const y1 = positions[from].y;
        const x2 = positions[to].x;
        const y2 = positions[to].y;

        svg.append('line')
          .attr('x1', x1).attr('y1', y1)
          .attr('x2', x2).attr('y2', y2)
          .attr('stroke', '#475569')
          .attr('stroke-width', 2)
          .attr('marker-end', directed ? 'url(#arrow)' : '');

        const mx = (x1 + x2) / 2;
        const my = (y1 + y2) / 2;
        svg.append('text')
          .attr('x', mx).attr('y', my - 6)
          .attr('text-anchor', 'middle')
          .attr('fill', '#94a3b8')
          .attr('font-size', '12px')
          .text(weight);
      }
    }

    if (directed) {
      svg.append('defs').append('marker')
        .attr('id', 'arrow')
        .attr('viewBox', '0 -5 10 10')
        .attr('refX', 32)
        .attr('refY', 0)
        .attr('markerWidth', 6)
        .attr('markerHeight', 6)
        .attr('orient', 'auto')
        .append('path')
        .attr('d', 'M0,-5L10,0L0,5')
        .attr('fill', '#475569');
    }

    for (const n of nodes) {
      const g = svg.append('g')
        .attr('transform', `translate(${positions[n].x}, ${positions[n].y})`)
        .style('cursor', 'grab')
        .on('dblclick', () => onremovenode(n));

      g.append('circle')
        .attr('r', radius)
        .attr('fill', nodeColor(n))
        .attr('stroke', '#0f172a')
        .attr('stroke-width', 3);

      g.append('text')
        .attr('text-anchor', 'middle')
        .attr('dominant-baseline', 'central')
        .attr('fill', 'white')
        .attr('font-size', '16px')
        .attr('font-weight', 'bold')
        .attr('pointer-events', 'none')
        .text(n);

      const drag = d3.drag<SVGGElement, unknown>()
        .on('start', function () {
          d3.select(this).style('cursor', 'grabbing');
        })
        .on('drag', function (event) {
          positions[n] = {
            x: Math.max(radius, Math.min(width - radius, event.x)),
            y: Math.max(radius, Math.min(height - radius, event.y))
          };
          draw();
        })
        .on('end', function () {
          d3.select(this).style('cursor', 'grab');
        });

      g.call(drag as any);
    }

    const legend = [
      { color: '#334155', label: 'Unvisited' },
      { color: '#60a5fa', label: 'Frontier' },
      { color: '#f59e0b', label: 'Current' },
      { color: '#22c55e', label: 'Visited' },
    ];
    legend.forEach((item, i) => {
      const lx = 20;
      const ly = height - 100 + i * 22;
      svg.append('circle').attr('cx', lx).attr('cy', ly).attr('r', 7).attr('fill', item.color);
      svg.append('text').attr('x', lx + 16).attr('y', ly + 1)
        .attr('dominant-baseline', 'central')
        .attr('fill', '#94a3b8').attr('font-size', '12px').text(item.label);
    });
    if (stepInfo) {
    svg.append('rect')
        .attr('x', width - 160).attr('y', 12)
        .attr('width', 148).attr('height', 32)
        .attr('rx', 8)

    svg.append('text')
        .attr('x', width - 86).attr('y', 32)
        .attr('text-anchor', 'middle')
        .attr('fill', '#fff')
        .attr('font-size', '16px')
        .text(stepInfo);
    }
  }

  $effect(() => {
    nodes; edges; visited; frontier; currentNode; directed; width; height;
    if (svgEl) draw();
  });
</script>

<div class="canvas-container" bind:this={containerEl}>
  <svg
    bind:this={svgEl}
    width={width}
    height={height}
    style="display:block;"
  />
</div>

<style>
  .canvas-container {
    width: 100%;
    height: 100%;
  }
</style>