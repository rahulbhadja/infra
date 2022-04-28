(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[203],{1431:function(e,t,n){"use strict";n.d(t,{Z:function(){return o}});var r=n(5893);function o(e){var t=e.iconPath,n=e.size,o=void 0===n?8:n,i=e.position;return(0,r.jsx)("div",{className:"flex items-center justify-center bg-gradient-to-br from-violet-400/30 to-pink-200/30 rounded-full ".concat("center"===i?"mx-auto my-4":""),children:(0,r.jsx)("div",{className:"flex bg-black items-center justify-center rounded-full w-16 h-16 m-0.5",children:(0,r.jsx)("img",{className:"w-".concat(o," h-").concat(o),src:t})})})}},9540:function(e,t,n){"use strict";n.d(t,{Z:function(){return h}});var r=n(4051),o=n.n(r),i=n(5893),a=n(1664),l=n.n(a),s=n(1163),c=n(8100),u=n(661);function f(e,t){(null==t||t>e.length)&&(t=e.length);for(var n=0,r=new Array(t);n<t;n++)r[n]=e[n];return r}function d(e,t,n,r,o,i,a){try{var l=e[i](a),s=l.value}catch(c){return void n(c)}l.done?t(s):Promise.resolve(s).then(r,o)}function p(e){return function(){var t=this,n=arguments;return new Promise((function(r,o){var i=e.apply(t,n);function a(e){d(i,r,o,a,l,"next",e)}function l(e){d(i,r,o,a,l,"throw",e)}a(void 0)}))}}function m(e){return function(e){if(Array.isArray(e))return f(e)}(e)||function(e){if("undefined"!==typeof Symbol&&null!=e[Symbol.iterator]||null!=e["@@iterator"])return Array.from(e)}(e)||function(e,t){if(!e)return;if("string"===typeof e)return f(e,t);var n=Object.prototype.toString.call(e).slice(8,-1);"Object"===n&&e.constructor&&(n=e.constructor.name);if("Map"===n||"Set"===n)return Array.from(n);if("Arguments"===n||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n))return f(e,t)}(e)||function(){throw new TypeError("Invalid attempt to spread non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function h(e){var t,n=e.children,r=(0,s.useRouter)(),a=(0,c.ZP)("/v1/identities/self").data,f=(0,c.ZP)("/v1/version").data,d=(0,u.A)(),h=d.admin,v=d.loading,x=(0,c.kY)().mutate;if(v)return null;function y(){return(y=p(o().mark((function e(){return o().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:fetch("/v1/logout",{method:"POST"}),x("/v1/identities/self",void 0),r.replace("/login");case 3:case"end":return e.stop()}}),e)})))).apply(this,arguments)}var g=[{name:"Clusters",href:"/destinations",icon:"/destinations.svg"},{name:"Identity Providers",href:"/providers",icon:"/providers.svg",admin:!0}],b=[{name:"Settings",href:"/settings",icon:"/settings.svg",admin:!0}],j=!0,w=!1,N=void 0;try{for(var k,C=m(g).concat(m(b))[Symbol.iterator]();!(j=(k=C.next()).done);j=!0){var P=k.value;if(r.asPath.startsWith(P.href)&&P.admin&&!h)return r.replace("/"),null}}catch(A){w=!0,N=A}finally{try{j||null==C.return||C.return()}finally{if(w)throw N}}return(0,i.jsxs)("div",{className:"flex h-full relative",children:[(0,i.jsxs)("nav",{className:"flex-none flex w-64 lg:w-72 flex-col inset-y-0 px-2 overflow-y-auto",children:[(0,i.jsx)("div",{className:"flex-shrink-0 flex items-center my-12 lg:my-18 px-6 select-none",children:(0,i.jsx)(l(),{href:"/",children:(0,i.jsx)("a",{children:(0,i.jsx)("img",{className:"h-[18px] w-auto",src:"infra.svg",alt:"Infra"})})})}),(0,i.jsx)("div",{className:"flex-1 space-y-1.5 px-3 select-none",children:g.map((function(e){return(0,i.jsx)(l(),{href:e.href,children:(0,i.jsxs)("a",{href:e.href,className:"\n                  ".concat(r.asPath.startsWith(e.href)?"bg-purple-200/10 text-white":"text-gray-500 hover:bg-purple-200/5 hover:text-gray-300","\n                  rounded-lg py-2 px-3 flex items-center text-sm font-medium transition-colors duration-100\n                  ").concat(e.admin&&!h?"opacity-30 pointer-events-none":"","\n                "),children:[(0,i.jsx)("img",{src:e.icon,className:"".concat(r.asPath.startsWith(e.href)?"":"opacity-30"," mr-3 flex-shrink-0 h-5 w-5")}),e.name]})},e.name)}))}),(0,i.jsxs)("div",{className:"relative group mx-2 my-5 h-16 hover:h-[178px] hover:bg-purple-100/5 transition-height transition-size px-4 duration-300 ease-in-out rounded-xl overflow-hidden",children:[(0,i.jsxs)("div",{className:"flex items-center space-x-4 mt-4 mb-2",children:[(0,i.jsx)("div",{className:"bg-purple-100/10 flex-none flex items-center justify-center w-9 h-9 py-1.5 rounded-lg capitalize font-bold select-none",children:null===a||void 0===a||null===(t=a.name)||void 0===t?void 0:t[0]}),(0,i.jsxs)("div",{children:[(0,i.jsx)("div",{className:"text-gray-300 text-sm font-medium overflow-hidden overflow-ellipsis leading-none",children:null===a||void 0===a?void 0:a.name}),h&&(0,i.jsx)("div",{className:"text-gray-400 text-xs leading-none my-1 capitalize",children:"Admin"})]})]}),(0,i.jsxs)("div",{className:"w-full px-2 py-1 items-center opacity-0 group-hover:opacity-100 transition-opacity duration-300 select-none text-sm",children:[(0,i.jsxs)("div",{onClick:function(){return function(){return y.apply(this,arguments)}()},className:"w-full flex items-center opacity-50 hover:opacity-75 py-2 cursor-pointer",children:[(0,i.jsx)("img",{src:"/signout.svg",className:"opacity-50 group-hover:opacity-75 h-3 mr-3"}),(0,i.jsx)("div",{className:"text-purple-50/40 group-hover:text-purple-50",children:"Logout"})]}),b.map((function(e){return(0,i.jsx)(l(),{href:e.href,children:(0,i.jsxs)("a",{className:"w-full flex -ml-1 opacity-50 hover:opacity-75 py-2 ".concat(e.admin&&!h?"pointer-events-none opacity-20":""),children:[(0,i.jsx)("img",{src:e.icon,className:"opacity-50 group-hover:opacity-75 mr-3 w-5 h-5"}),(0,i.jsx)("div",{className:"text-purple-50/40 group-hover:text-purple-50",children:e.name})]})},e.name)}))]}),(0,i.jsxs)("div",{className:"px-2 pt-1 pb-3 text-xs text-purple-50/30",children:["Infra version ",null===f||void 0===f?void 0:f.version]})]})]}),(0,i.jsx)("main",{className:"w-full mx-auto xl:max-w-4xl 2xl:max-w-5xl overflow-x-hidden overflow-y-scroll",children:n})]})}},1399:function(e,t,n){"use strict";n.d(t,{Z:function(){return l}});var r=n(5893),o=n(7294),i=n(7533),a=n(8771);function l(e){var t=e.open,n=e.setOpen,l=e.onSubmit,s=e.title,c=e.message,u=(0,o.useRef)(null);return(0,r.jsx)(i.V,{as:"div",className:"fixed z-10 inset-0 overflow-y-auto",initialFocus:u,open:t,onClose:function(){return n(!1)},children:(0,r.jsxs)("div",{className:"flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0",children:[(0,r.jsx)(i.V.Overlay,{className:"fixed inset-0 bg-black bg-opacity-75 transition-opacity"}),(0,r.jsx)("span",{className:"hidden sm:inline-block sm:align-middle sm:h-screen","aria-hidden":"true",children:"\u200b"}),(0,r.jsx)("div",{className:"relative inline-block bg-gradient-to-tr from-[#B06363] to-[#FF00C7] rounded-3xl text-left overflow-hidden shadow-xl transform transition-all my-8 align-middle max-w-2xl w-full p-px",children:(0,r.jsxs)("div",{className:"bg-black px-10 pt-12 pb-8 rounded-3xl",children:[(0,r.jsxs)("div",{className:"flex items-start",children:[(0,r.jsx)("div",{className:"rounded-full bg-gradient-to-tr from-[#B06363] to-[#FF00C7]",children:(0,r.jsx)("div",{className:"flex h-14 w-14 items-center justify-center rounded-full bg-black m-0.5",children:(0,r.jsx)(a.Z,{className:"h-6 w-6 text-[#D3398F]","aria-hidden":"true"})})}),(0,r.jsxs)("div",{className:"mt-1 ml-5 text-left",children:[(0,r.jsx)(i.V.Title,{as:"h3",className:"text-base leading-6 font-bold text-white",children:s}),(0,r.jsx)("p",{className:"text-sm text-gray-400 my-0.5",children:c})]})]}),(0,r.jsxs)("div",{className:"mt-8 text-sm flex flex-row-reverse",children:[(0,r.jsx)("button",{type:"button",className:"w-auto inline-flex justify-center rounded-full bg-gradient-to-tr from-[#B06363] to-[#FF00C7] font-medium text-white focus:outline-none focus:ring-2 focus:ring-[#FF00C7] ml-3",onClick:function(){return l()},children:(0,r.jsx)("div",{className:"bg-black px-10 py-3.5 rounded-full m-0.5",children:"Delete"})}),(0,r.jsx)("button",{type:"button",className:"w-auto inline-flex items-center justify-center rounded-full px-10 py-3.5 bg-black hover:opacity-75 font-medium text-white focus:outline-none focus:ring-2 focus:ring-zinc-600",onClick:function(){return n(!1)},ref:u,children:"Cancel"})]})]})})]})})}},5857:function(e,t,n){"use strict";n.d(t,{Z:function(){return l}});var r=n(5893),o=n(9521);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{},r=Object.keys(n);"function"===typeof Object.getOwnPropertySymbols&&(r=r.concat(Object.getOwnPropertySymbols(n).filter((function(e){return Object.getOwnPropertyDescriptor(n,e).enumerable})))),r.forEach((function(t){i(e,t,n[t])}))}return e}function l(e){var t=e.columns,n=e.data,i=e.getRowProps,l=void 0===i?function(){return{}}:i,s=e.showHeader,c=void 0===s||s,u=(0,o.useTable)({columns:t,data:n}),f=u.getTableProps,d=u.getTableBodyProps,p=u.headerGroups,m=u.rows,h=u.prepareRow;return(0,r.jsxs)("table",a({className:"w-full table-auto"},f(),{children:[c&&(0,r.jsx)("thead",{className:"border-b border-zinc-800",children:p.map((function(e){return(0,r.jsx)("tr",a({},e.getHeaderGroupProps(),{children:e.headers.map((function(e){return(0,r.jsx)("th",a({className:"text-left uppercase py-2 font-normal text-sm text-gray-400"},e.getHeaderProps(),{children:e.render("Header")}),e.id)}))}),e.id)}))}),(0,r.jsx)("tbody",a({},d(),{children:m.map((function(e){return h(e),(0,r.jsx)("tr",a({className:"text-sm group"},e.getRowProps(l(e)),{children:e.cells.map((function(e){return(0,r.jsx)("td",a({className:"py-1 group-first:pt-3"},e.getCellProps(),{children:e.render("Cell")}),e.id)}))}),e.id)}))}))]}))}},661:function(e,t,n){"use strict";n.d(t,{A:function(){return o}});var r=n(8100);function o(){var e=(0,r.ZP)("/v1/identities/self").data,t=(0,r.ZP)((function(){return"/v1/identities/".concat(null===e||void 0===e?void 0:e.id,"/grants?resource=infra")})),n=t.data,o=t.error;return{loading:!n&&!o,admin:!!(null===n||void 0===n?void 0:n.find((function(e){return"admin"===e.privilege})))}}},1551:function(e,t,n){"use strict";function r(e,t){(null==t||t>e.length)&&(t=e.length);for(var n=0,r=new Array(t);n<t;n++)r[n]=e[n];return r}function o(e,t){return function(e){if(Array.isArray(e))return e}(e)||function(e,t){var n=null==e?null:"undefined"!==typeof Symbol&&e[Symbol.iterator]||e["@@iterator"];if(null!=n){var r,o,i=[],a=!0,l=!1;try{for(n=n.call(e);!(a=(r=n.next()).done)&&(i.push(r.value),!t||i.length!==t);a=!0);}catch(s){l=!0,o=s}finally{try{a||null==n.return||n.return()}finally{if(l)throw o}}return i}}(e,t)||function(e,t){if(!e)return;if("string"===typeof e)return r(e,t);var n=Object.prototype.toString.call(e).slice(8,-1);"Object"===n&&e.constructor&&(n=e.constructor.name);if("Map"===n||"Set"===n)return Array.from(n);if("Arguments"===n||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n))return r(e,t)}(e,t)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var i,a=(i=n(7294))&&i.__esModule?i:{default:i},l=n(1003),s=n(880),c=n(9246);var u={};function f(e,t,n,r){if(e&&l.isLocalURL(t)){e.prefetch(t,n,r).catch((function(e){0}));var o=r&&"undefined"!==typeof r.locale?r.locale:e&&e.locale;u[t+"%"+n+(o?"%"+o:"")]=!0}}var d=function(e){var t,n=!1!==e.prefetch,r=s.useRouter(),i=a.default.useMemo((function(){var t=o(l.resolveHref(r,e.href,!0),2),n=t[0],i=t[1];return{href:n,as:e.as?l.resolveHref(r,e.as):i||n}}),[r,e.href,e.as]),d=i.href,p=i.as,m=a.default.useRef(d),h=a.default.useRef(p),v=e.children,x=e.replace,y=e.shallow,g=e.scroll,b=e.locale;"string"===typeof v&&(v=a.default.createElement("a",null,v));var j=(t=a.default.Children.only(v))&&"object"===typeof t&&t.ref,w=o(c.useIntersection({rootMargin:"200px"}),3),N=w[0],k=w[1],C=w[2],P=a.default.useCallback((function(e){h.current===p&&m.current===d||(C(),h.current=p,m.current=d),N(e),j&&("function"===typeof j?j(e):"object"===typeof j&&(j.current=e))}),[p,j,d,C,N]);a.default.useEffect((function(){var e=k&&n&&l.isLocalURL(d),t="undefined"!==typeof b?b:r&&r.locale,o=u[d+"%"+p+(t?"%"+t:"")];e&&!o&&f(r,d,p,{locale:t})}),[p,d,k,b,n,r]);var A={ref:P,onClick:function(e){t.props&&"function"===typeof t.props.onClick&&t.props.onClick(e),e.defaultPrevented||function(e,t,n,r,o,i,a,s){("A"!==e.currentTarget.nodeName.toUpperCase()||!function(e){var t=e.currentTarget.target;return t&&"_self"!==t||e.metaKey||e.ctrlKey||e.shiftKey||e.altKey||e.nativeEvent&&2===e.nativeEvent.which}(e)&&l.isLocalURL(n))&&(e.preventDefault(),t[o?"replace":"push"](n,r,{shallow:i,locale:s,scroll:a}))}(e,r,d,p,x,y,g,b)},onMouseEnter:function(e){t.props&&"function"===typeof t.props.onMouseEnter&&t.props.onMouseEnter(e),l.isLocalURL(d)&&f(r,d,p,{priority:!0})}};if(e.passHref||"a"===t.type&&!("href"in t.props)){var O="undefined"!==typeof b?b:r&&r.locale,S=r&&r.isLocaleDomain&&l.getDomainLocale(p,O,r&&r.locales,r&&r.domainLocales);A.href=S||l.addBasePath(l.addLocale(p,O,r&&r.defaultLocale))}return a.default.cloneElement(t,A)};t.default=d,("function"===typeof t.default||"object"===typeof t.default&&null!==t.default)&&(Object.assign(t.default,t),e.exports=t.default)},9246:function(e,t,n){"use strict";function r(e,t){(null==t||t>e.length)&&(t=e.length);for(var n=0,r=new Array(t);n<t;n++)r[n]=e[n];return r}function o(e,t){return function(e){if(Array.isArray(e))return e}(e)||function(e,t){var n=null==e?null:"undefined"!==typeof Symbol&&e[Symbol.iterator]||e["@@iterator"];if(null!=n){var r,o,i=[],a=!0,l=!1;try{for(n=n.call(e);!(a=(r=n.next()).done)&&(i.push(r.value),!t||i.length!==t);a=!0);}catch(s){l=!0,o=s}finally{try{a||null==n.return||n.return()}finally{if(l)throw o}}return i}}(e,t)||function(e,t){if(!e)return;if("string"===typeof e)return r(e,t);var n=Object.prototype.toString.call(e).slice(8,-1);"Object"===n&&e.constructor&&(n=e.constructor.name);if("Map"===n||"Set"===n)return Array.from(n);if("Arguments"===n||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(n))return r(e,t)}(e,t)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}Object.defineProperty(t,"__esModule",{value:!0}),t.useIntersection=function(e){var t=e.rootRef,n=e.rootMargin,r=e.disabled||!l,u=i.useRef(),f=o(i.useState(!1),2),d=f[0],p=f[1],m=o(i.useState(t?t.current:null),2),h=m[0],v=m[1],x=i.useCallback((function(e){u.current&&(u.current(),u.current=void 0),r||d||e&&e.tagName&&(u.current=function(e,t,n){var r=function(e){var t,n={root:e.root||null,margin:e.rootMargin||""},r=c.find((function(e){return e.root===n.root&&e.margin===n.margin}));r?t=s.get(r):(t=s.get(n),c.push(n));if(t)return t;var o=new Map,i=new IntersectionObserver((function(e){e.forEach((function(e){var t=o.get(e.target),n=e.isIntersecting||e.intersectionRatio>0;t&&n&&t(n)}))}),e);return s.set(n,t={id:n,observer:i,elements:o}),t}(n),o=r.id,i=r.observer,a=r.elements;return a.set(e,t),i.observe(e),function(){if(a.delete(e),i.unobserve(e),0===a.size){i.disconnect(),s.delete(o);var t=c.findIndex((function(e){return e.root===o.root&&e.margin===o.margin}));t>-1&&c.splice(t,1)}}}(e,(function(e){return e&&p(e)}),{root:h,rootMargin:n}))}),[r,h,n,d]),y=i.useCallback((function(){p(!1)}),[]);return i.useEffect((function(){if(!l&&!d){var e=a.requestIdleCallback((function(){return p(!0)}));return function(){return a.cancelIdleCallback(e)}}}),[d]),i.useEffect((function(){t&&v(t.current)}),[t]),[x,d,y]};var i=n(7294),a=n(4686),l="undefined"!==typeof IntersectionObserver;var s=new Map,c=[];("function"===typeof t.default||"object"===typeof t.default&&null!==t.default)&&(Object.assign(t.default,t),e.exports=t.default)},1664:function(e,t,n){e.exports=n(1551)}}]);