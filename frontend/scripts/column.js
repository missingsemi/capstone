class Column extends HTMLElement {
  color;

  constructor() {
    super();

    this.attachShadow({mode: 'open'});

    this.color = [
      '255,0,0',
      '0,255,0',
      '0,0,255',
      '255,255,0',
      '255,0,255',
      '0,255,255',
    ][Math.floor(Math.random() * 6)]

    const style = document.createElement('style');
    style.textContent = `
      h2 {
        line-height: 2rem;
        margin-top: 0;
        margin-bottom: 10px;
        font-size: 2rem;
        border-bottom: 2px solid white;
        white-space: nowrap;
      }

      .column-list {
        height: calc(calc(100% - 2rem) - 12px);
        display: grid;
        grid-template-rows: repeat(38, calc(100% / 38));
        grid-auto-columns: 1fr;
        list-style: none;
        padding: 0;
        margin: 0;
      }

      .session {
        height: 100%;
        width: 100%;
      }
    `
    
    const titleH2 = document.createElement('h2');
    titleH2.setAttribute('class', 'title');

    const columnUl = document.createElement('ul');
    columnUl.setAttribute('class', 'column-list');

    this.shadowRoot.appendChild(style);
    this.shadowRoot.appendChild(titleH2);
    this.shadowRoot.appendChild(columnUl);
  }

  static get observedAttributes() {
    return [
      'column:title',
      'column:sessions',
    ];
  }

  connectedCallback() {
    this.updateTitle(this.getAttribute('column:title'));
    this.updateSessions('[]', this.getAttribute('column:session'));
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (name === 'column:title') {
      this.updateTitle(newValue);
    }

    if (name === 'column:sessions') {
      console.log(newValue);
      this.updateSessions(oldValue, newValue);
    }
  }

  updateTitle(title) {
    const titleH2 = this.shadowRoot.querySelector('.title');
    titleH2.textContent = title;
  }

  updateSessions(oldValue, newValue) {
    oldValue ??= '[]';
    newValue ??= '[]';

    const prev = JSON.parse(oldValue);
    const curr = JSON.parse(newValue);
    const columnList = this.shadowRoot.querySelector('.column-list');
    
    // Finding the added and deleted elements so we don't have to update the entire column.
    const currIds = new Set(curr.map(session => session.id));
    const prevIds = prev.map(session => session.id);

    const prevSet = new Set(prevIds);
    const added = curr.filter(session => !prevSet.has(session.id));

    const currSet = new Set(currIds);
    const deletedIds = prevIds.filter(id => !currSet.has(id));

    for (const id of deletedIds) {
      columnList.removeChild(columnList.getElementById(`${id}`));
    }

    for (const session of added) {
      const cs = document.createElement('calendar-session');
      const rowSpan = Math.floor(session.duration / 15);
      const startTime = new Date(session.time);
      const rowStart = 4 * (startTime.getHours() - 8) + Math.floor(startTime.getMinutes() / 15) + 1;
      cs.style = `grid-row: ${rowStart} / span ${rowSpan};`;
      cs.setAttribute('class', 'session');
      cs.setAttribute('session:duration', `${session.duration}`);
      cs.setAttribute('session:time', session.time);
      cs.setAttribute('session:name', session.username);
      cs.setAttribute('session:color', this.color);
      cs.id = `${session.id}`;
      columnList.appendChild(cs);
    }
  }
}

customElements.define('calendar-column', Column);