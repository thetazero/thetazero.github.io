customElements.define('theta-nav', class extends HTMLElement {
  constructor() {
    super()
    const shadow = this.attachShadow({ mode: 'open' })
    shadow.innerHTML = `
    <a class='logo' href='https://thetazero.github.io'>
      <span>0</span>
      <span>Θ</span>
    </div>
    <style>
      :host{
        position:fixed;
        left:0px;
        top:0px;
        width:100%;
        background:var(--background_nav);
        --size:50px;
        height:var(--size);
        z-index:100;
      }

      @media only screen and (max-width: 768px) {
        
      }

      .logo{
        width:var(--size);
        height:var(--size);
        cursor:pointer;
        transition:color var(--text_transition_speed);
        color:var(--text);
        display:inline-block;
      }

      span{
        position:absolute;
        transform:translate(-50%,-50%);
        display:inline-block;
        user-select:none;
        font-weight:100;
        color:inherit;
        font-size:var(--size);
        top:calc(var(--size) / 2);
        left:calc(var(--size) / 2);
        line-height:var(--size);
      }
      .logo:hover{
        color:var(--text_accent) !important;
      }
    </style>`
  }
})