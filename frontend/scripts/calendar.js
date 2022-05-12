class Calendar extends HTMLElement {
  columns = {};

  constructor() {
    super();

    this.attachShadow({mode: 'open'});

    const style = document.createElement('style');
    style.textContent = `
      .calendar {
        height: 100%;
        min-width: 100%;
        width: min-content;
        background-color: #111115;
        color: white;
        display: flex;
        gap: 20px;
      }

      .scroll {
        height: 100%;
        width: 100%;
        display: flex;
        gap: 20px;
      }

      .scale {
        position: sticky;
        left: 0;
        top: 0;
        background-color: #111115;
        padding-right: 10px;
      }

      .column {
        height: 100%;
      }

      .column>h2 {
        line-height: 2rem;
        margin-top: 0;
        margin-bottom: 10px;
        font-size: 2rem;
      }

      .column-list {
        height: calc(calc(100% - 2rem) - 12px);
        display: grid;
        grid-template-rows: repeat(19, calc(100% / 19));
        margin: 0;
        list-style: none;
        padding: 0;
      }

      .column-list>li {
        justify-self: end;
      }
    `;

    const calendarDiv = document.createElement('div');
    calendarDiv.setAttribute('class', 'calendar');

    const scaleDiv = document.createElement('div');
    scaleDiv.setAttribute('class', 'column scale');
    const scaleH2 = document.createElement('h2');
    scaleH2.textContent = 'Time';
    const scaleUl = document.createElement('ul');
    scaleUl.setAttribute('class', 'column-list');
    this.createScale(scaleUl);

    const scrollDiv = document.createElement('div');
    scrollDiv.setAttribute('class', 'scroll');

    scaleDiv.appendChild(scaleH2);
    scaleDiv.appendChild(scaleUl);
    calendarDiv.appendChild(scaleDiv);
    calendarDiv.appendChild(scrollDiv);
    this.shadowRoot.appendChild(style);
    this.shadowRoot.appendChild(calendarDiv);
  }

  static get observedAttributes() {
    return [];
  }

  createScale(scaleUl) {
    const times = [
      '8:00',
      '8:30',
      '9:00',
      '9:30',
      '10:00',
      '10:30',
      '11:00',
      '11:30',
      '12:00',
      '12:30',
      '1:00',
      '1:30',
      '2:00',
      '2:30',
      '3:00',
      '3:30',
      '4:00',
      '4:30',
      '5:00',
    ];
    for (const time of times) {
      const li = document.createElement('li');
      li.textContent = time;
      scaleUl.appendChild(li);
    }
  }

  attributeChangedCallback() {}

  connectedCallback() {
    if (this.isConnected) this.update();
  }
  
  async update() {
    const machinesResp = await fetch('/api/machines');
    const machinesJson = await machinesResp.json();

    const scrollBox = this.shadowRoot.querySelector('.scroll');

    for (const machine of machinesJson) {
      if (!(machine.id in this.columns)) {
        let column = document.createElement('calendar-column');
        column.setAttribute('class', 'column')
        scrollBox.appendChild(column);
        this.columns[machine.id] = column;
      }

      this.columns[machine.id].setAttribute('column:title', machine.titleName);
    }

    const sessionsResp = await fetch('/api/schedule');
    const sessionsJson = await sessionsResp.json();

    const sessions = {};
    for (const session of sessionsJson) {
      if (session.machine in sessions) {
        sessions[session.machine].push(session);
      } else {
        sessions[session.machine] = [session];
      }
    }

    for (const machineId in this.columns) {
      this.columns[machineId].setAttribute('column:sessions', JSON.stringify(sessions[machineId] ?? []));
    }
  }
}

customElements.define('calendar-calendar', Calendar);