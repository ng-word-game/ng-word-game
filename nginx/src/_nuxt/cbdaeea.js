(window.webpackJsonp=window.webpackJsonp||[]).push([[8,3],{377:function(t,e,n){"use strict";n.r(e);n(1),n(13),n(141),n(17),n(92);var o=n(47),r=n(110),c=n.n(r),l=n(113),d=n(77),v=Object(o.a)({name:"OdaiCom",setup:function(){var t=Object(o.j)().app,e=Object(o.c)(l.b),n=Object(o.h)(""),r=Object(o.h)(!1),v=Object(o.k)(),f=Object(o.h)(""),h=Object(o.h)(!1);if(!e)throw new Error("store undefined");var data=Object(o.g)(e.data);e.state.socket||v.push({name:"index"});var m=Object(o.h)([]);Object(o.d)((function(){f.value=e.data.thema}));var _=function(s){return new Promise((function(e){c.a.post("https://labs.goo.ne.jp/api/hiragana",{app_id:t.$config.apiKey,sentence:s,output_type:"hiragana"}).then((function(t){e(t.data.converted)}))}))};Object(o.l)(data,(function(){data.game_state===d.b.GameStart&&(h.value=!1,console.log("start"),v.push({name:"play"}))}),{deep:!0}),Object(o.l)(n,(function(){r.value=null==n.value.match(/^[ぁ-んー　]{4,6}$/),""!==n.value&&c.a.get("".concat(location.protocol,"//").concat(location.host,"/api/suggest/search?hl=ja&q=").concat(n.value,"&output=toolbar"),{responseType:"document"}).then((function(t){var e=t.data.querySelectorAll("suggestion"),n=[];m.value=[];for(var i=0;i<e.length;i++)null!=e[i].getAttribute("data")&&null===e[i].getAttribute("data").match(/.* .*/)&&n.push(e[i].getAttribute("data"));n=n.sort((function(a,b){return a.length-b.length}));for(var o=0;o<n.length&&!(o>3);o++)_(n[o]).then((function(t){return m.value.push(t)}))}))}));return{wordRef:n,wordValid:r,thema:f,registered:function(){e.setMyWord(n.value),h.value=!0,e.state.socket&&e.state.socket.send(JSON.stringify({type:d.a.SetWord,word:n.value}))},waiting:h,suggestionsRef:m}}}),f=n(76),component=Object(f.a)(v,(function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"d-flex justify-content-center align-items-center",staticStyle:{height:"100vh"}},[n("div",{staticClass:"card",staticStyle:{width:"50%"}},[n("div",{staticClass:"card-body"},[n("h5",{staticClass:"card-title text-center"},[t._v("\n        お題\n      ")]),t._v(" "),n("h1",{staticClass:"my-4 text-center"},[t._v("\n        "+t._s(t.thema)+"\n      ")]),t._v(" "),n("div",{staticClass:"form-group d-flex flex-column justify-content-center mt-4"},[n("span",{staticClass:"form-text"},[t._v("4~6文字でひらがなのみ")]),t._v(" "),n("input",{directives:[{name:"model",rawName:"v-model",value:t.wordRef,expression:"wordRef"}],staticClass:"form-control",attrs:{type:"text",placeholder:"ことばをかく"},domProps:{value:t.wordRef},on:{input:function(e){e.target.composing||(t.wordRef=e.target.value)}}}),t._v(" "),t.wordValid?n("div",{staticClass:"form-text",staticStyle:{color:"red"}},[t._v("4~6文字でひらがなのみを使用してください")]):t._e(),t._v(" "),n("p",{staticClass:"text-center mt-3"},[t._v("候補一覧")]),t._v(" "),t._l(t.suggestionsRef,(function(e,o){return n("div",{key:o,staticClass:"border mx-auto text-center",staticStyle:{width:"50%"}},[n("div",[t._v(t._s(e))])])})),t._v(" "),n("div",{staticClass:"text-center"},[t.suggestionsRef?n("a",{attrs:{href:"http://www.goo.ne.jp/"}},[n("img",{staticStyle:{width:"5%"},attrs:{src:"//u.xgoo.jp/img/sgoo.png",alt:"supported by goo",title:"supported by goo"}})]):t._e()]),t._v(" "),n("div",{staticClass:"mx-auto mt-3"},[!t.waiting&&t.wordValid?n("button",{staticClass:"btn btn-outline-info",attrs:{disabled:"",type:"submit"},on:{click:t.registered}},[t._v("\n            決定\n          ")]):t.waiting?n("p",[t._v("\n            他のプレイヤーを待っています...\n          ")]):n("button",{staticClass:"btn btn-outline-info",attrs:{type:"submit"},on:{click:t.registered}},[t._v("\n            決定\n          ")])])],2)])])])}),[],!1,null,null,null);e.default=component.exports},379:function(t,e,n){"use strict";n.r(e);var o=n(47),r=n(377),c=Object(o.a)({name:"OdaiPage",components:{OdaiCom:r.default},setup:function(){return{}}}),l=n(76),component=Object(l.a)(c,(function(){var t=this.$createElement;return(this._self._c||t)("OdaiCom")}),[],!1,null,null,null);e.default=component.exports}}]);