(window.webpackJsonp=window.webpackJsonp||[]).push([[10,6],{347:function(t,e,n){"use strict";n.r(e);var l=n(48),c=n(111),r=n(90),o=Object(l.a)({name:"ResultCom",setup:function(){var t=Object(l.c)(c.b);if(!t)throw new Error("store undefined");var e=Object(l.k)();return{store:t,resultMode:function(){return t.data.game_state===r.b.GameEnd},stopMode:function(){return t.data.game_state===r.b.GameStop},goIndex:function(){return e.push({name:"index"})}}}}),d=n(75),component=Object(d.a)(o,(function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"d-flex justify-content-center align-items-center",staticStyle:{height:"100vh"}},[n("div",{staticClass:"card",staticStyle:{width:"50%"}},[n("div",{staticClass:"card-body"},[t.resultMode()?n("h3",{staticClass:"card-title text-center"},[t._v("\n        Result\n      ")]):t.stopMode()?n("h3",{staticClass:"card-title text-center"},[t._v("\n        対戦相手が退出しました\n      ")]):n("h3",{staticClass:"card-title text-center"},[t._v("\n        接続が切断されました\n      ")]),t._v(" "),t.resultMode()?n("h4",{staticClass:"text-center"},[t._v("\n        "+t._s(t.store.data.winner)+"の勝利\n      ")]):t._e(),t._v(" "),n("div",{staticClass:"d-flex flex-column justify-content-center"},[n("div",{staticClass:"mx-auto mt-3"},[n("button",{staticClass:"btn btn-outline-info",attrs:{type:"submit"},on:{click:t.goIndex}},[t._v("\n            最初に戻る\n          ")])])])])])])}),[],!1,null,null,null);e.default=component.exports},352:function(t,e,n){"use strict";n.r(e);var l=n(11),c=n(347),r=l.default.extend({name:"ResultPage",components:{Result:c.default}}),o=n(75),component=Object(o.a)(r,(function(){var t=this.$createElement;return(this._self._c||t)("result")}),[],!1,null,null,null);e.default=component.exports;installComponents(component,{Result:n(347).default})}}]);