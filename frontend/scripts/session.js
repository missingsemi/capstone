class Session extends HTMLElement {
  constructor() {
    super();

    this.root = this.attachShadow({mode: 'open'});

    const style = document.createElement('style');
    style.textContent = `
      .session {
        height: 100%;
        width: 100%;
        padding-left: 10px;
        padding-right: 10px;
        border: solid #111115 1px;
        border-top-left-radius: 5px;
        border-bottom-left-radius: 5px;
        overflow: hidden;
        box-sizing: border-box;
      }

      .name {
        font-size: 1.5rem;
        font-weight: bold;
        line-height: 1.5rem;
        margin-top: 5px;
        margin-bottom: 5px;
        box-sizing: border-box;
      }
    `

    const sessionDiv = document.createElement('div');
    sessionDiv.setAttribute('class', 'session');
    const nameH3 = document.createElement('h3');
    nameH3.setAttribute('class', 'name');
    const timeSpan = document.createElement('span');
    timeSpan.setAttribute('class', 'time');
    timeSpan.append(document.createTextNode('Time'));
    timeSpan.append(document.createTextNode(' for '));
    timeSpan.append(document.createTextNode('duration'));
    timeSpan.append(document.createTextNode(' minutes.'));

    sessionDiv.appendChild(nameH3);
    sessionDiv.appendChild(timeSpan);

    this.root.appendChild(style);
    this.root.appendChild(sessionDiv);
  }

  updateName(name) {
    this.root.querySelector('.name').textContent = name;
  }

  updateTime(time) {
    const timeFormatter = new Intl.DateTimeFormat('en-US', {
      timeStyle: 'short'
    });
    const timeString = timeFormatter.format(new Date(time))
    this.shadowRoot.querySelector('.time').childNodes[0].textContent = timeString;
  }

  updateDuration(duration) {
    this.shadowRoot.querySelector('.time').childNodes[2].textContent = duration;
  }

  updateColor(color) {
    this.shadowRoot.querySelector('.session').style = `background-color: rgba(${color},0.1); border-left: 5px solid rgb(${color});`;
  }

  connectedCallback() {
    if (this.isConnected) {
      this.updateColor(this.getAttribute('session:color'));
      this.updateName(this.getAttribute('session:name'));
      this.updateTime(this.getAttribute('session:time'));
      this.updateDuration(this.getAttribute('session:duration'));
    }
  }

  attributeChangedCallback(name, oldValue, newValue) {
    if (this.isConnected) {
      if (name === 'session:name') {
        this.updateName(newValue);
      }
      if (name === 'session:time') {
        this.updateTime(newValue);
      }
      if (name === 'session:duration') {
        this.updateDuration(newValue);
      }
      if (name === 'session:color') {
        this.updateColor(newValue);
      }
    }
  }

  parseProps() {
    this.props.duration = Number(this.getAttribute('session:duration'));
    this.props.time = this.getAttribute('session:time');
    this.props.name = this.getAttribute('session:name');
  }

  static get observedAttributes() {
    return [
      'session:duration',
      'session:time',
      'session:name',
    ];
  }
}

customElements.define('calendar-session', Session);