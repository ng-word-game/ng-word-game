(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{111:function(e,t,n){"use strict";n.d(t,"b",(function(){return d}));n(28),n(4),n(60),n(1);var r=n(8),o=n(90),c=Object(r.l)({game_state:o.b.Initial,next_turn:"",result:0,thema:"",users:[],winner:"",word_state:{},words:{"":""},ng_chars:[{name:"",char:""}]}),l=Object(r.l)({name:"",myword:"",socket:null});t.a={state:Object(r.m)(l),data:Object(r.m)(c),setName:function(e){l.name=e},setMyWord:function(e){l.myword=e},setSocket:function(s){l.socket=s},setData:function(e){Object.assign(c,e)}};var d=Symbol("key")},226:function(e,t,n){var content=n(328);content.__esModule&&(content=content.default),"string"==typeof content&&(content=[[e.i,content,""]]),content.locals&&(e.exports=content.locals);(0,n(78).default)("d829a5ca",content,!0,{sourceMap:!1})},239:function(e,t,n){"use strict";var r=n(45),o=n(111),c=n(90),l=n(29),d=(n(73),n(1),function(e){var t=new FileReader;return new Promise((function(n,r){t.onerror=function(){t.abort(),r()},t.onload=function(){n(t.result)},t.readAsText(e)}))}),f=function(){var e=Object(l.a)(regeneratorRuntime.mark((function e(t){var n;return regeneratorRuntime.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,d(t);case 2:return n=e.sent,e.abrupt("return",JSON.parse(n));case 4:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),m=Object(r.a)({setup:function(){Object(r.e)(o.b,o.a);var e=Object(r.c)(o.b);if(!e)throw new Error("store undefined");var t=Object(r.g)(e.state),data=Object(r.f)(e.data),n=Object(r.g)(0),l=Object(r.i)();return Object(r.j)(t,(function(){t.value.socket&&0===n.value&&(t.value.socket.addEventListener("message",(function(t){f(t.data).then((function(t){console.log(t),e.setData(t)}))})),n.value=1)}),{deep:!0}),Object(r.j)(data,(function(){data.game_state!==c.b.GameStop&&data.game_state!==c.b.GameEnd||l.push({name:"result"})}),{deep:!0}),{}}}),v=(n(327),n(76)),component=Object(v.a)(m,(function(){var e=this.$createElement,t=this._self._c||e;return t("transition",{attrs:{name:"page"}},[t("nuxt")],1)}),[],!1,null,null,null);t.a=component.exports},241:function(e,t,n){n(242),n(243),e.exports=n(259)},277:function(e,t,n){var content=n(278);content.__esModule&&(content=content.default),"string"==typeof content&&(content=[[e.i,content,""]]),content.locals&&(e.exports=content.locals);(0,n(78).default)("77cb1d58",content,!0,{sourceMap:!1})},278:function(e,t,n){var r=n(77)(!1);r.push([e.i,".page-enter-active,.page-leave-active{transition:opacity .5s}.page-enter,.page-leave-to{opacity:0}",""]),e.exports=r},327:function(e,t,n){"use strict";n(226)},328:function(e,t,n){var r=n(77)(!1);r.push([e.i,'html{font-family:"Source Sans Pro",-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif;font-size:16px;word-spacing:1px;-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%;-moz-osx-font-smoothing:grayscale;-webkit-font-smoothing:antialiased;box-sizing:border-box}*,:after,:before{box-sizing:border-box;margin:0}',""]),e.exports=r},90:function(e,t,n){"use strict";var r,o,c;n.d(t,"b",(function(){return r})),n.d(t,"a",(function(){return o})),n.d(t,"c",(function(){return c})),function(e){e[e.Initial=0]="Initial",e[e.PlayerFilled=1]="PlayerFilled",e[e.GameStart=2]="GameStart",e[e.GameEnd=3]="GameEnd",e[e.GameStop=4]="GameStop"}(r||(r={})),function(e){e[e.SetWord=0]="SetWord",e[e.SetNgChar=1]="SetNgChar"}(o||(o={})),function(e){e[e.HiddenWord=0]="HiddenWord",e[e.OpenedWord=1]="OpenedWord"}(c||(c={}))}},[[241,11,1,12]]]);