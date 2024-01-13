<script>
  import {
    Table,
    TableBody,
    TableBodyCell,
    TableBodyRow,
    TableHead,
    TableHeadCell,
  } from "flowbite-svelte";
  export let data;
</script>

<Table hoverable={true}>
  <TableHead>
    <TableHeadCell>Key</TableHeadCell>
    <TableHeadCell>Value</TableHeadCell>
  </TableHead>
  <TableBody>
    {#each data as item}
      <TableBodyRow>
        <TableBodyCell>{item.key}</TableBodyCell>
        <TableBodyCell>{item.value}</TableBodyCell>
      </TableBodyRow>
    {/each}
  </TableBody>
</Table>
