<script>
  import { onMount } from "svelte";

/*
  Id       int       `json:"id"`
	UserId   string    `json:"userId"`
	GroupIds []string  `json:"groupIds"`
	Machine  string    `json:"machine"`
	Reason   string    `json:"reason"`
	Duration int64     `json:"duration"`
	Time     time.Time `json:"time"`
*/

  /**
   * Represents a session for a machine.
   * @typedef {Object} Session
   * @property {number} id
   * @property {string} userId
   * @property {string[]} groupIds
   * @property {string} machine
   * @property {string} reason
   * @property {number} duration
   * @property {string} time
   */

   /**
    * @type {Session[]}
    */
  let sessions = [];
  onMount(async () => {
    const resp = await fetch("/api/schedule");
    sessions = await resp.json()
  })
</script>

{#each sessions as s}
  <h1>{s.userId}</h1>
{/each}