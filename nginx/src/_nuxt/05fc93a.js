(window.webpackJsonp=window.webpackJsonp||[]).push([[5,4],{371:function(t,e,n){var content=n(373);content.__esModule&&(content=content.default),"string"==typeof content&&(content=[[t.i,content,""]]),content.locals&&(t.exports=content.locals);(0,n(79).default)("44d54db6",content,!0,{sourceMap:!1})},372:function(t,e,n){"use strict";n(371)},373:function(t,e,n){var r=n(78)(!1);r.push([t.i,".rect[data-v-17463148]{border:1.2px solid;display:table-cell;text-align:center;vertical-align:middle}.rect p[data-v-17463148]{margin:10px;font-size:1.5rem;font-weight:700}.opened[data-v-17463148]{background-color:grey}.hide[data-v-17463148]{background-color:#000}",""]),t.exports=r},374:function(t,e,n){"use strict";n.r(e);var r=n(47),c=Object(r.a)({name:"SquareCom",props:{charaInfo:{type:Array,default:null}},setup:function(t){return{props:t}}}),o=(n(372),n(76)),component=Object(o.a)(c,(function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("div",{staticClass:"d-flex"},t._l(t.props.charaInfo,(function(e,r){return n("div",{key:r,staticClass:"rect mx-2",class:[e.isOpen?"opened":"",e.isHide?"hide":""]},[e.isHide?n("p",[t._v("\n        ・\n      ")]):n("p",[t._v("\n        "+t._s(e.char)+"\n      ")])])})),0)])}),[],!1,null,"17463148",null);e.default=component.exports},378:function(t,e,n){"use strict";n.r(e);n(2),n(1),n(42),n(13),n(141),n(3);var r=n(47),c=n(113),o=n(77),l=n(374),d=Object(r.a)({name:"PlayCom",components:{SquareCom:l.default},setup:function(){var t=Object(r.c)(c.b);if(!t)throw new Error("store undefined");var e=Object(r.k)();t.state.socket||e.push({name:"index"});var n=t.data.users.filter((function(u){return u.Id===t.state.clientId}))[0],l=t.data.users.filter((function(u){return u.Id!==t.state.clientId})),d=Object(r.h)(""),v=Object(r.h)(!0),h=Object(r.h)(),data=Object(r.g)(t.data),f=Object(r.h)(t.data.next_turn),_=Object(r.h)(),m=Object(r.g)([]),C=Object(r.h)(t.data.thema),x=Object(r.h)(t.data.ng_chars),I=Object(r.h)(0),y=function(){h.value.hide(),t.state.socket?(t.state.socket.send(JSON.stringify({type:o.a.SetNgChar,ng_char:d.value})),d.value=""):console.log("websocket is not found")},O=Object(r.h)(),j=function(){I.value=60;var t=setInterval((function(){if(I.value<1)return y(),void clearInterval(t);I.value-=1}),1e3);O.value=t};Object(r.d)((function(){w(),j()})),Object(r.l)(data,(function(){w(),f.value=data.next_turn,f.value===n.Id&&(clearInterval(O.value),j()),x.value=data.ng_chars,console.log(x)}),{deep:!0}),Object(r.l)(d,(function(){v.value=null==d.value.match(/^[ぁ-んー　]{1}$/)}));var w=function(){var e=[];t.data.word_state[n.Id].forEach((function(t){for(var n in t)e.push({char:n,isOpen:t[n]===o.c.OpenedWord,isHide:!1})})),_.value=e,l.forEach((function(e){var n=[];t.data.word_state[e.Id].forEach((function(t){for(var e in t)n.push({char:e,isOpen:t[e]===o.c.OpenedWord,isHide:t[e]===o.c.HiddenWord})})),0===m.filter((function(t){return t.id===e.Id})).length?m.push({id:e.Id,chars:n}):m.filter((function(t){return t.id===e.Id}))[0].chars=n}))};return{user:n,userCharInfo:_,anotherUsers:l,anotherUsersCharInfo:m,nextTurn:f,ngCharas:x,ngChar:d,ngCharValid:v,modalRef:h,registerNgChar:y,thema:C,timer:I}}}),v=n(76),component=Object(v.a)(d,(function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"d-flex justify-content-center align-items-center flex-column",staticStyle:{height:"100vh"}},[t._l(t.anotherUsers,(function(e){return n("div",{staticClass:"mb-3"},[n("p",{staticClass:"text-center",staticStyle:{"font-size":"1.2rem","font-weight":"bold"}},[t._v("\n      "+t._s(e.Name)+"のワード\n    ")]),t._v(" "),e&&t.anotherUsersCharInfo!==[]&&t.anotherUsersCharInfo.filter((function(t){return t.id===e.Id}))!==[]&&t.anotherUsersCharInfo.filter((function(t){return t.id===e.Id}))[0]&&t.anotherUsersCharInfo.filter((function(t){return t.id===e.Id}))[0].chars?n("SquareCom",{attrs:{"chara-info":t.anotherUsersCharInfo.filter((function(t){return t.id===e.Id}))[0].chars}}):t._e()],1)})),t._v(" "),n("div",{staticClass:"d-flex align-items-center justify-content-around",staticStyle:{height:"40vh",width:"80%"}},[n("div",{staticStyle:{width:"30%"}},[n("div",{staticClass:"card"},[n("p",{staticClass:"text-center",staticStyle:{"font-size":"100%"}},[t._v("\n          お題\n        ")]),t._v(" "),n("p",{staticClass:"text-center",staticStyle:{"font-size":"200%","font-weight":"bold"}},[t._v("\n          "+t._s(t.thema)+"\n        ")])]),t._v(" "),n("div",{staticClass:"my-2 d-flex justify-content-center"},[t.nextTurn===t.user.Id?n("b-button",{directives:[{name:"b-modal",rawName:"v-b-modal.modal-1",modifiers:{"modal-1":!0}}],staticStyle:{width:"80%"},attrs:{variant:"outline-info"}},[t._v("\n          文字を選ぶ\n        ")]):n("p",[t._v("\n          相手を待っています...\n        ")])],1),t._v(" "),t.nextTurn===t.user.Id?n("p",{staticClass:"text-center"},[t._v("制限時間")]):t._e(),t._v(" "),t.nextTurn===t.user.Id?n("p",{staticClass:"text-center"},[t._v("残り"),n("span",{staticClass:"font-weight-bold",staticStyle:{color:"red"}},[t._v(t._s(t.timer))]),t._v("秒")]):t._e()]),t._v(" "),n("div",{staticClass:"card",staticStyle:{height:"100%",width:"60%"}},[n("div",{staticClass:"card-body"},[n("h5",{staticClass:"text-center"},[t._v("\n          NGリスト\n        ")]),t._v(" "),n("div",{staticClass:"table-responsive",staticStyle:{height:"80%"}},[n("table",{staticClass:"table table-sm"},[t._m(0),t._v(" "),t._l(t.ngCharas,(function(e){return n("tbody",{key:e.name+e.char},[n("tr",[n("td",{staticClass:"text-center"},[t._v("\n                  "+t._s(e.name)+"\n                ")]),t._v(" "),""===e.char?n("td",{staticClass:"text-center"},[t._v("\n                  時間切れ\n                ")]):n("td",{staticClass:"text-center"},[t._v("\n                  "+t._s(e.char)+"\n                ")])])])}))],2)])])])]),t._v(" "),n("div",{staticClass:"mt-3"},[n("p",{staticClass:"text-center",staticStyle:{"font-size":"1.2rem","font-weight":"bold"}},[t._v("\n      "+t._s(t.user.Name)+"のワード\n    ")]),t._v(" "),t.userCharInfo?n("SquareCom",{attrs:{"chara-info":t.userCharInfo}}):t._e()],1),t._v(" "),n("b-modal",{ref:"modalRef",attrs:{id:"modal-1","hide-footer":""}},[n("div",{staticClass:"d-block text-center"},[n("h3",[t._v("NG文字を入力しよう")])]),t._v(" "),n("input",{directives:[{name:"model",rawName:"v-model",value:t.ngChar,expression:"ngChar"}],staticClass:"form-control",attrs:{type:"text",placeholder:""},domProps:{value:t.ngChar},on:{input:function(e){e.target.composing||(t.ngChar=e.target.value)}}}),t._v(" "),t.ngCharValid?n("div",{staticClass:"form-text",staticStyle:{color:"red"}},[t._v("ひらがな1文字のみ入力できます")]):t._e(),t._v(" "),n("b-button",{staticClass:"mt-2",attrs:{disabled:t.ngCharValid,variant:"outline-info",block:""},on:{click:t.registerNgChar}},[t._v("\n      決定\n    ")])],1)],2)}),[function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("thead",[n("tr",[n("th",{staticClass:"text-center",attrs:{scope:"col"}},[t._v("\n                  名前\n                ")]),t._v(" "),n("th",{staticClass:"text-center",attrs:{scope:"col"}},[t._v("\n                  NG文字\n                ")])])])}],!1,null,null,null);e.default=component.exports}}]);