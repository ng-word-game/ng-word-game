(window.webpackJsonp=window.webpackJsonp||[]).push([[7,2],{346:function(e,t,n){"use strict";n.r(t);n(14),n(31),n(1),n(42),n(20),n(17);var r=n(47),o=n(111),c=n(76),l=n(30),d=(n(79),function(e){var t=new FileReader;return new Promise((function(n,r){t.onerror=function(){t.abort(),r()},t.onload=function(){n(t.result)},t.readAsText(e)}))}),f=function(){var e=Object(l.a)(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,d(t);case 2:return n=e.sent,e.abrupt("return",JSON.parse(n));case 4:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),x=Object(r.a)({name:"PlayerJoin",setup:function(){var e=Object(r.j)().app,t=Object(r.c)(o.b),n=Object(r.h)(""),l=Object(r.k)();if(!t)throw new Error("store undefined");var data=Object(r.g)(t.data),d=Object(r.h)(!1),x=Object(r.h)(!1),v=function(){for(var e="xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".split(""),i=0,t=e.length;i<t;i++)switch(e[i]){case"x":e[i]=Math.floor(16*Math.random()).toString(16);break;case"y":e[i]=(Math.floor(4*Math.random())+8).toString(16)}return e.join("")}();function m(){console.log("close"),l.push({name:"result"})}Object(r.l)(data,(function(){console.log("here"),data.game_state===c.b.PlayerFilled&&(d.value=!1,l.push({name:"odai"}))}),{deep:!0});return{name:n,connectError:x,registered:function(){t.setName(n.value),t.setClientId(v),d.value=!0;var r=new WebSocket("".concat(e.$config.wsURL,"/?id=").concat(v,"&name=").concat(n.value));console.log(r),r.addEventListener("error",(function(){x.value=!0})),r.addEventListener("open",(function(){console.log("open"),r.addEventListener("message",(function(e){f(e.data).then((function(e){console.log(e),t.setData(e)}))})),r.addEventListener("close",m,!1),t.setSocket(r)}))},waiting:d}}}),v=n(75),component=Object(v.a)(x,(function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"d-flex justify-content-center align-items-center",staticStyle:{height:"100vh"}},[n("div",{staticClass:"card",staticStyle:{width:"50%"}},[n("div",{staticClass:"card-body"},[n("h5",{staticClass:"card-title text-center"},[e._v("\n        NG単語ゲーム\n      ")]),e._v(" "),n("div",{staticClass:"form-group d-flex flex-column justify-content-center"},[n("input",{directives:[{name:"model",rawName:"v-model",value:e.name,expression:"name"}],staticClass:"form-control",attrs:{type:"email",placeholder:"ユーザー名"},domProps:{value:e.name},on:{input:function(t){t.target.composing||(e.name=t.target.value)}}}),e._v(" "),n("div",{staticClass:"mx-auto mt-3"},[e.waiting?e.connectError?e._e():n("p",[e._v("\n            参加者を待っています....\n          ")]):n("button",{staticClass:"btn btn-outline-info",attrs:{type:"submit",disabled:e.waiting},on:{click:e.registered}},[e._v("\n            参加する\n          ")]),e._v(" "),e.connectError?n("p",[e._v("サーバーと接続できません")]):e._e()])])])])])}),[],!1,null,null,null);t.default=component.exports},353:function(e,t,n){"use strict";n.r(t);var r=n(11),o=n(346),c=r.default.extend({name:"IndexPage",components:{Join:o.default}}),l=n(75),component=Object(l.a)(c,(function(){var e=this.$createElement;return(this._self._c||e)("join")}),[],!1,null,null,null);t.default=component.exports;installComponents(component,{Join:n(346).default})}}]);