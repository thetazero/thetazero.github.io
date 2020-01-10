customElements.define('theta-module', class extends HTMLElement {
  constructor() {
    super()
    const shadow = this.attachShadow({ mode: 'open' })
    console.log(this.title)
    shadow.innerHTML = `
    <a class='link'>
    <span class='text'>
      <span class='title'></span><br>
      <span class='desc'></span>
    </span>
    <style>
      .text{
        position:relative;
        top:50%;
        display:inline-block;
        transform:translateY(-50%);
      }
      .title{
        font-size:2em;
      }
      .desc{
        font-size:0.8em;
      }
      @media only screen and (max-width: 768px) {
        .title{
          font-size:2em;
        }
      }
      :host{
        background:var(--background_emphasize);
        padding:5px;
        border-radius:5px;
        cursor:pointer;
        transition:color 100ms ease-in-out;
        vertical-align:top;
      }
      .link{
        color:var(--text_emphasize);
        width:100%;
        height:100%;
        display:block;
      }
      :host(:hover) .title{
        color:var(--text_accent);
      }
      </a>
    </style>`
  }
  connectedCallback() {
    this.shadowRoot.querySelector('span.title').innerText = this.title
    this.shadowRoot.querySelector('span.desc').innerText = this.desc
    this.shadowRoot.querySelector('a.link').setAttribute('href', this.getAttribute('href'))
  }
  get title() {
    return this.getAttribute('title')
  }
  get desc() {
    return this.getAttribute('desc')
  }
})