(window.webpackJsonp=window.webpackJsonp||[]).push([[1],{114:function(e,t,n){"use strict";n.d(t,"b",(function(){return d}));n(28),n(4),n(59),n(1);var o=n(8),r=n(71),c=Object(o.m)({game_state:r.b.Initial,next_turn:"",turn:0,result:0,thema:"",users:null,winner:"",word_state:{},words:{"":""},ng_chars:[{name:"",char:""}]}),l=Object(o.m)({name:"",clientId:"",myword:"",socket:null});t.a={state:Object(o.n)(l),data:Object(o.n)(c),setName:function(e){l.name=e},setClientId:function(e){l.clientId=e},setMyWord:function(e){l.myword=e},setSocket:function(s){l.socket=s},setData:function(e){Object.assign(c,e)}};var d=Symbol("key")},229:function(e,t,n){var content=n(339);content.__esModule&&(content=content.default),"string"==typeof content&&(content=[[e.i,content,""]]),content.locals&&(e.exports=content.locals);(0,n(79).default)("d829a5ca",content,!0,{sourceMap:!1})},250:function(e,t,n){"use strict";n(42);var o=n(47),r=n(114),c=n(71),l=Object(o.b)({setup:function(){Object(o.g)(r.b,r.a);var e=Object(o.d)(r.b);if(!e)throw new Error("store undefined");var t=Object(o.i)(e.state),data=Object(o.h)(e.data),n=!0,l=function(){return setInterval((function(){null!==e.state.socket&&e.state.socket.send(JSON.stringify({type:c.a.PING}))}),1e4)};return Object(o.m)(data,(function(){null!==e.state.socket?data.game_state!==c.b.GameStop&&data.game_state!==c.b.GameEnd&&data.game_state!==c.b.GameDrawEnd||e.state.socket.close():clearInterval(l())}),{deep:!0}),Object(o.f)((function(){n&&(n=!1,console.log("here"),l())})),Object(o.e)((function(){t.value.socket&&t.value.socket.close()})),{}}}),d=(n(338),n(77)),component=Object(d.a)(l,(function(){var e=this.$createElement,t=this._self._c||e;return t("transition",{attrs:{name:"page"}},[t("nuxt")],1)}),[],!1,null,null,null);t.a=component.exports},252:function(e,t,n){n(253),n(254),e.exports=n(257)},288:function(e,t,n){var content=n(289);content.__esModule&&(content=content.default),"string"==typeof content&&(content=[[e.i,content,""]]),content.locals&&(e.exports=content.locals);(0,n(79).default)("77cb1d58",content,!0,{sourceMap:!1})},289:function(e,t,n){var o=n(78)(!1);o.push([e.i,".page-enter-active,.page-leave-active{transition:opacity .5s}.page-enter,.page-leave-to{opacity:0}",""]),e.exports=o},338:function(e,t,n){"use strict";n(229)},339:function(e,t,n){var o=n(78)(!1);o.push([e.i,'html{font-family:"Source Sans Pro",-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif;font-size:16px;word-spacing:1px;-ms-text-size-adjust:100%;-webkit-text-size-adjust:100%;-moz-osx-font-smoothing:grayscale;-webkit-font-smoothing:antialiased;box-sizing:border-box}*,:after,:before{box-sizing:border-box;margin:0}',""]),e.exports=o},71:function(e,t,n){"use strict";var o,r,c;n.d(t,"b",(function(){return o})),n.d(t,"a",(function(){return r})),n.d(t,"c",(function(){return c})),function(e){e[e.Initial=0]="Initial",e[e.PlayerFilled=1]="PlayerFilled",e[e.GameStart=2]="GameStart",e[e.GameEnd=3]="GameEnd",e[e.GameStop=4]="GameStop",e[e.GameDrawEnd=5]="GameDrawEnd"}(o||(o={})),function(e){e[e.SetWord=0]="SetWord",e[e.SetNgChar=1]="SetNgChar",e[e.PING=2]="PING"}(r||(r={})),function(e){e[e.HiddenWord=0]="HiddenWord",e[e.OpenedWord=1]="OpenedWord"}(c||(c={}))}},[[252,13,2,14]]]);