(function(t){function e(e){for(var s,n,i=e[0],c=e[1],l=e[2],u=0,d=[];u<i.length;u++)n=i[u],Object.prototype.hasOwnProperty.call(r,n)&&r[n]&&d.push(r[n][0]),r[n]=0;for(s in c)Object.prototype.hasOwnProperty.call(c,s)&&(t[s]=c[s]);p&&p(e);while(d.length)d.shift()();return a.push.apply(a,l||[]),o()}function o(){for(var t,e=0;e<a.length;e++){for(var o=a[e],s=!0,n=1;n<o.length;n++){var c=o[n];0!==r[c]&&(s=!1)}s&&(a.splice(e--,1),t=i(i.s=o[0]))}return t}var s={},r={app:0},a=[];function n(t){return i.p+"js/"+({about:"about"}[t]||t)+"."+{about:"ef8669ba"}[t]+".js"}function i(e){if(s[e])return s[e].exports;var o=s[e]={i:e,l:!1,exports:{}};return t[e].call(o.exports,o,o.exports,i),o.l=!0,o.exports}i.e=function(t){var e=[],o=r[t];if(0!==o)if(o)e.push(o[2]);else{var s=new Promise((function(e,s){o=r[t]=[e,s]}));e.push(o[2]=s);var a,c=document.createElement("script");c.charset="utf-8",c.timeout=120,i.nc&&c.setAttribute("nonce",i.nc),c.src=n(t);var l=new Error;a=function(e){c.onerror=c.onload=null,clearTimeout(u);var o=r[t];if(0!==o){if(o){var s=e&&("load"===e.type?"missing":e.type),a=e&&e.target&&e.target.src;l.message="Loading chunk "+t+" failed.\n("+s+": "+a+")",l.name="ChunkLoadError",l.type=s,l.request=a,o[1](l)}r[t]=void 0}};var u=setTimeout((function(){a({type:"timeout",target:c})}),12e4);c.onerror=c.onload=a,document.head.appendChild(c)}return Promise.all(e)},i.m=t,i.c=s,i.d=function(t,e,o){i.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:o})},i.r=function(t){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},i.t=function(t,e){if(1&e&&(t=i(t)),8&e)return t;if(4&e&&"object"===typeof t&&t&&t.__esModule)return t;var o=Object.create(null);if(i.r(o),Object.defineProperty(o,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var s in t)i.d(o,s,function(e){return t[e]}.bind(null,s));return o},i.n=function(t){var e=t&&t.__esModule?function(){return t["default"]}:function(){return t};return i.d(e,"a",e),e},i.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},i.p="/",i.oe=function(t){throw console.error(t),t};var c=window["webpackJsonp"]=window["webpackJsonp"]||[],l=c.push.bind(c);c.push=e,c=c.slice();for(var u=0;u<c.length;u++)e(c[u]);var p=l;a.push([0,"chunk-vendors"]),o()})({0:function(t,e,o){t.exports=o("56d7")},2162:function(t,e,o){},"56d7":function(t,e,o){"use strict";o.r(e);o("e260"),o("e6cf"),o("cca6"),o("a79d");var s=o("2b0e"),r=function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("v-app",[o("v-app-bar",{attrs:{app:"","clipped-left":"","clipped-right":""}},[o("v-app-bar-nav-icon",{attrs:{disabled:"",_click:"drawer = !drawer"}}),o("v-toolbar-title",[t._v("Dashboard")]),o("v-spacer"),o("v-badge",t._b({attrs:{content:t.alertCount,dot:t.hideAlertBadge,color:"error",left:"",_overlap:"","offset-x":"23","offset-y":"23"}},"v-badge",t.alertBadgeIconProp,!1),[o("v-btn",{attrs:{icon:""},on:{click:t.openAlertSheet}},[o("v-icon",[t._v("mdi-bell")])],1)],1),o("v-tooltip",{attrs:{bottom:"","open-delay":1e3},scopedSlots:t._u([{key:"activator",fn:function(e){var s=e.on,r=e.attrs;return[o("v-btn",t._g(t._b({attrs:{icon:""},on:{click:t.changeDarkMode}},"v-btn",r,!1),s),[o("v-icon",[t._v("mdi-theme-light-dark")])],1)]}}])},[o("span",[t._v("Switch to day/night mode")])])],1),o("v-main",[t.socketConnected?t._e():o("v-banner",{attrs:{app:"","single-line":"",_color:"white",icon:"mdi-wifi-strength-alert-outline","icon-color":"red"},scopedSlots:t._u([{key:"actions",fn:function(){return[o("v-btn",{attrs:{text:"",color:"primary"},on:{click:t.reconnectSocket}},[t._v("Reconnect")])]},proxy:!0}],null,!1,4246916530)},[t._v(" Lost socket connection to API Server ")]),o("v-container",{attrs:{fluid:""}},[t.skeletonLoading?o("v-row",t._l(9,(function(t){return o("v-col",{key:t,attrs:{cols:"12",lg:"3"}},[o("v-skeleton-loader",{staticClass:"mx-auto",attrs:{_loading:"skeletonLoading",transition:"scale-transition","max-width":"300",type:"card"}})],1)})),1):0==t.hostsReport.length?o("v-row",[o("v-col",[t.apiError?t.socketConnected?o("v-alert",{attrs:{text:"",color:"info",border:"left"}},[o("div",[t._v("Please waiting for the first report being aggregated at CentMonit...")])]):o("v-alert",{attrs:{text:"",color:"error",border:"left"}},[o("h3",{staticClass:"headline"},[t._v(t._s(t.apiErrorMessage))]),o("div",[t._v(" Is CentMonit daemon process running normally on the server side (at address "+t._s(""+this.$GCONFIG.api_base_url)+")?"),o("br"),t._v(" Try restart it then reload your browser (Ctrl + F5). ")])]):o("v-alert",{attrs:{text:"",color:"info",border:"left"}},[o("h3",{staticClass:"headline"},[t._v("No data found")]),o("div",[t._v(" You may need to wait for CentMonit to receive first report data from Monit agents."),o("br"),t._v(" If you've waited so long, then check your Monit log (/var/log/monit.log) to see whether it successfully send report data to CentMonit?"),o("br"),t._v(" Please follow the setup manual from installation folder to verify Monit and CentMonit's configuration. ")])])],1)],1):o("div",[o("v-row",{staticClass:"justify-center"},[o("small",{class:[t.lastReportTimeStyleClass],staticStyle:{color:"#2196F3"}},[t._v(" Good hosts: "+t._s(t.totalGoodHosts())+"/"+t._s(Object.keys(t.hostsState).length)+" - Last report: "+t._s(t.lastHostsReportTime.toLocaleString())+" ")])]),o("v-row",t._l(t.hostsReport,(function(e,s){return o("v-col",{key:s+"-"+e.alertMessage,attrs:{_key:"index",cols:"12",_sm:"6",md:"6",lg:"4"}},[o("HostCard",{attrs:{report:e},on:{hostStateChanged:t.onHostStateChanged}})],1)})),1)],1),o("v-bottom-sheet",{staticStyle:{"overflow-y":"scroll"},model:{value:t.alertSheet,callback:function(e){t.alertSheet=e},expression:"alertSheet"}},[o("v-card",[o("v-card-title",{attrs:{_class:"justify-center"}},[o("span",{staticClass:"text-subtitle-1"},[t._v("Realtime Channel")]),o("v-icon",[t._v("mdi-circle-small")]),o("span",{class:[t.socketStatusStyleClass]},[t._v(t._s(t.socketConnectionStatusDesc))]),o("v-spacer"),o("v-tooltip",{attrs:{bottom:"","open-delay":1e3},scopedSlots:t._u([{key:"activator",fn:function(e){var s=e.on,r=e.attrs;return[o("v-btn",t._g(t._b({attrs:{icon:""},on:{click:t.alertReset}},"v-btn",r,!1),s),[o("v-icon",[t._v("mdi-notification-clear-all")])],1)]}}])},[o("span",[t._v("Clear all messages (old messages will be removed automatically)")])]),o("v-btn",{attrs:{icon:""},on:{click:function(e){t.alertSheet=!t.alertSheet}}},[o("v-icon",[t._v("mdi-chevron-down")])],1)],1),o("v-divider"),o("v-card-text",{staticClass:"pt-4",staticStyle:{height:"80vh","overflow-y":"scroll"}},[t.socketConnected?t._e():o("div",{staticClass:"text-center"},[o("v-btn",{staticClass:"mb-2",attrs:{_loading:"btnSocketReconnectLoading",_disabled:"btnSocketReconnectLoading",outlined:"",color:"primary"},on:{click:t.reconnectSocket}},[t._v(" Reconnect "),o("v-icon",{attrs:{right:""}},[t._v("mdi-sync")])],1)],1),t._l(t.alert_messages,(function(e,s){return o("v-alert",{key:s,attrs:{border:"left","colored-border":"",color:t.getAlertColor(e.type),elevation:"1"}},[Boolean(e.host)?o("span",{staticClass:"text-subtile-1 font-weight-bold"},[t._v(" "+t._s(e.host)+" "),Boolean(e.service)?[o("v-icon",{attrs:{small:""}},[t._v("mdi-circle-small")]),t._v(" "+t._s(e.service)+" ")]:t._e(),o("br")],2):t._e(),o("span",{staticClass:"text-caption"},[o("v-icon",{attrs:{small:""}},[t._v("mdi-calendar-clock")]),t._v(" "+t._s(e.time.toLocaleString())),o("br")],1),Boolean(e.serviceTypeDesc)?o("span",{staticClass:"text-caption"},[o("v-icon",{attrs:{small:""}},[t._v("mdi-tag-text-outline")]),t._v(" "+t._s(e.serviceTypeDesc)),o("br")],1):t._e(),o("div",{staticClass:"mt-1"},[t._v(" "+t._s(e.message)+" ")])])}))],2)],1)],1)],1)],1),o("v-footer",{attrs:{app:"",_class:"grey _dark_en-3 white--text"}},[o("div",{staticClass:"text-caption",attrs:{_class:"ml-3 text-caption text-sm-body-2"}},[t._v(" Made with "),o("v-icon",{staticClass:"red--text",staticStyle:{"vertical-align":"middle"},attrs:{small:""}},[t._v("mdi-heart-outline")]),t._v(" by "),o("a",{attrs:{_class:"white--text",href:"https://vuetifyjs.com",target:"_blank"}},[t._v("Vuetify")])],1),o("v-spacer"),o("div",{staticClass:"font-italic caption",attrs:{_class:"mr-3 hidden-xs-only"}},[t._v(" Version "+t._s(t.APP_VERSION)+" ")])],1)],1)},a=[],n=(o("4de4"),o("4160"),o("b64b"),o("159b"),o("5530")),i=function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("v-card",{staticClass:"mx-auto",attrs:{"max-width":"350","min-height":"206",_click:"test"}},[o("v-list-item",{attrs:{"_three-line":""}},[o("v-list-item-avatar",[o("v-icon",{class:"mdi-loading"===t.hostLedDesc[1]?"mdi-spin":"",attrs:{color:t.hostLedDesc[0]}},[t._v(t._s(t.hostLedDesc[1]))])],1),o("v-list-item-content",[o("v-list-item-title",{staticClass:"text-subtile-1 font-weight-bold"},[t._v(" "+t._s(t.report.hostname)+" ")]),o("v-list-item-subtitle",{staticClass:"text-caption"},[o("span",[t._v(t._s(t.uptimeDesc)+" monit up")]),o("v-icon",{attrs:{small:""}},[t._v("mdi-circle-small")]),o("span",[t._v(t._s(t.pollTimeDesc)+" poll-cycle")])],1)],1)],1),o("v-card-text",[t.report.alertMessage?o("v-row",{staticClass:"text-center"},[o("v-col",{attrs:{cols:"12"}},[o("div",[t._v(t._s(t.report.alertMessage))])])],1):o("v-row",{staticClass:"text-center"},[o("v-col",{attrs:{cols:"4"}},[o("div",[o("span",{staticClass:"text-h4"},[t._v(t._s(t.report.goodServices))]),o("span",{staticClass:"text-caption"},[t._v(" /"+t._s(t.report.services))])]),o("div",[t._v("OK")])]),o("v-col",{attrs:{cols:"4"}},[o("div",{staticClass:"text-h4"},[t._v(t._s(t.report.failServices))]),o("div",[t._v("Fail")])]),o("v-col",{attrs:{cols:"4"}},[o("div",{staticClass:"text-h4"},[t._v(t._s(t.report.skipServices))]),o("div",[t._v("Skip")])])],1)],1),o("v-card-actions",[o("div",[o("span",{class:[t.cpuTextColor]},[t._v(t._s(t.cpuDesc)+" cpu")]),o("v-icon",{attrs:{small:""}},[t._v("mdi-circle-small")]),o("span",{class:[t.ramTextColor]},[t._v(t._s(t.ramDesc)+" ram")])],1),o("v-spacer"),o("v-tooltip",{attrs:{color:t.hostLedDesc[0],top:""},scopedSlots:t._u([{key:"activator",fn:function(e){var s=e.on,r=e.attrs;return[o("v-btn",t._g(t._b({attrs:{icon:""}},"v-btn",r,!1),s),[o("v-icon",[t._v("mdi-table")])],1)]}}]),model:{value:t.tooltipModel,callback:function(e){t.tooltipModel=e},expression:"tooltipModel"}},[o("v-simple-table",{scopedSlots:t._u([{key:"default",fn:function(){return[o("thead",[o("tr",[o("th",{staticClass:"text-left"},[t._v("Service")]),o("th",{staticClass:"text-left"},[t._v("Type")]),o("th",{staticClass:"text-left"},[t._v("Monitor")]),o("th",{staticClass:"text-left"},[t._v("Status")])])]),o("tbody",t._l(t.hostServices,(function(e){return o("tr",{key:t.report.id+"-"+e.name},[o("td",[t._v(t._s(e.name))]),o("td",[t._v(t._s(e.type))]),o("td",[t._v(t._s(t.getServiceMonitorDesc(e.monitor)))]),o("td",[o("v-icon",{attrs:{small:"",color:t.getServiceStatusColor(e.status)}},[t._v(" mdi-circle-medium ")])],1)])})),0)]},proxy:!0}])})],1)],1)],1)},c=[],l=(o("99af"),o("b680"),o("bc3a")),u=o.n(l),p={props:["report"],data:function(){return{tooltipModel:!1,hostServices:[]}},watch:{tooltipModel:function(t){t||this.__pre_fetch_services__()},report:{handler:function(t){t.alertMessage&&(this.hostServices=[])},deep:!0}},computed:{uptimeDesc:function(){var t=this.report.uptime;switch(!0){case t<60:return"".concat(t," secs");case t<3600:return"".concat((t/60).toFixed(1)," mins");case t<86400:return"".concat((t/3600).toFixed(1)," hrs");default:return"".concat((t/86400).toFixed(1)," days")}},pollTimeDesc:function(){var t=this.report.poll;switch(!0){case t<60:return"".concat(t," secs");case t<3600:return"".concat((t/60).toFixed(1)," mins");case t<86400:return"".concat((t/3600).toFixed(1)," hrs");default:return"".concat((t/86400).toFixed(1)," days")}},cpuDesc:function(){return this.report.cpu.toFixed(1)+"%"},ramDesc:function(){return this.report.ram.toFixed(1)+"%"},cpuTextColor:function(){var t="";switch(!0){case this.report.cpu<80:break;case this.report.cpu<95:t="amber--text";break;case this.report.cpu>=95:t="red--text";break}return t},ramTextColor:function(){var t="";switch(!0){case this.report.ram<80:break;case this.report.ram<95:t="amber--text";break;case this.report.ram>=95:t="red--text";break}return t},hostLedDesc:function(){var t="grey",e="mdi-loading",o="loading";return this.report.alertMessage?(t="red",e="mdi-message-alert"):this.report.services===this.report.failServices?(t="red",e="mdi-alert"):this.report.skipServices>=1||this.report.failServices>=1?(t="amber",e="mdi-alert"):this.report.services===this.report.goodServices&&(t="green",e="mdi-check-bold",o="good"),this.$emit("hostStateChanged",{id:this.report.id,status:o}),[t,e]}},created:function(){this.__pre_fetch_services__()},methods:{__pre_fetch_services__:function(){var t=this;if(!this.report.alertMessage){var e=this.report.id;u.a.get("".concat(this.$GCONFIG.api_base_url,"/api/hosts/").concat(e,"/report")).then((function(e){console.log("HostCard::__pre_fetch_services__() - get report response:",e),e.data.length>0&&(t.hostServices=[],e.data.forEach((function(e){t.hostServices.push(e)})))})).catch((function(t){console.log("HostCard::__pre_fetch_services__() - get report error:",t)}))}},getServiceStatusColor:function(t){switch(t){case 0:return"success";default:return"error"}},getServiceMonitorDesc:function(t){switch(t){case 1:return"Yes";case 0:return"No";case 2:return"Initializing...";default:return"UNKNOWN"}}}},d=p,h=o("2877"),v=o("6544"),m=o.n(v),f=o("8336"),_=o("b0af"),g=o("99d9"),b=o("62ad"),y=o("132d"),C=o("da13"),S=o("8270"),k=o("5d23"),x=o("0fd9"),w=o("1f4f"),T=o("2fa4"),V=o("3a2f"),R=Object(h["a"])(d,i,c,!1,null,null,null),A=R.exports;m()(R,{VBtn:f["a"],VCard:_["a"],VCardActions:g["a"],VCardText:g["b"],VCol:b["a"],VIcon:y["a"],VListItem:C["a"],VListItemAvatar:S["a"],VListItemContent:k["a"],VListItemSubtitle:k["b"],VListItemTitle:k["c"],VRow:x["a"],VSimpleTable:w["a"],VSpacer:T["a"],VTooltip:V["a"]});var M=o("2f62"),j={name:"App",components:{HostCard:A},data:function(){return{APP_VERSION:"2020.0-beta2",num:1,apiError:!1,apiErrorMessage:null,skeletonLoading:!0,drawer:!1,drawerRight:!1,alertSheet:!1,alertCount:0,alertEvents:[],items:[{title:"Dashboard",icon:"mdi-view-dashboard"},{title:"Photos",icon:"mdi-image"},{title:"About",icon:"mdi-help-box"}],socketConnection:null,socketConnected:!1,btnSocketReconnectLoading:!1,hostsReport:[],lastHostsReportTime:null,lastReportTimeStyleClass:"",hostsState:{},hostsNextReportTime:{}}},computed:Object(n["a"])(Object(n["a"])({},Object(M["b"])(["app_dark_mode","alert_messages"])),{},{hideAlertBadge:function(){return 0==this.alertCount&&this.socketConnected},alertBadgeIconProp:function(){return this.socketConnected?"":{icon:"mdi-close"}},socketConnectionStatusDesc:function(){return this.socketConnected?"connected":"not connect"},socketStatusStyleClass:function(){return this.socketConnected?"text-caption green--text":"text-caption red--text"}}),watch:{alertCount:function(t){document.title=0==t?"CentMonit":t<99?"(".concat(t,") CentMonit"):"(99+) CentMonit"}},created:function(){var t=this;console.log("App::created()..."),this.$vuetify.theme.dark=this.app_dark_mode,u.a.get("".concat(this.$GCONFIG.api_base_url,"/api/hosts/report")).then((function(e){console.log("App::created() - get report response:",e),t.hostsReport=e.data,t.lastHostsReportTime=new Date,t.hostsReport.forEach((function(e){e.alertMessage||t.$set(t.hostsNextReportTime,e.id,(new Date).getTime()+1e3*e.poll)}))})).catch((function(e){console.log("App::created() - get report error:",e),t.apiError=!0,t.apiErrorMessage=e.message})),setTimeout((function(){t.skeletonLoading=!1}),1500),this.__init_socket_connection__()},mounted:function(){var t=this;console.log("App:mounted()...");var e="bgJobID",o=localStorage.getItem(e);o&&(clearInterval(parseInt(o,10)),console.log("Cleared background job with id",o)),o=setInterval((function(){var e=t.hostsNextReportTime;console.warn((new Date).toLocaleTimeString()+" - hostsNextReportTime:",JSON.stringify(e)),Object.keys(e).forEach((function(o){var s=e[o],r=(new Date).getTime();r>s&&t.hostsReport.forEach((function(t){t.id===o&&(console.warn("No report from ",o,"-",t.hostname),t.alertMessage="No report within poll cycle")}))}))}),1e4),localStorage.setItem(e,o),console.log("Saved background job with id",o)},methods:{__highlight_report_time_text__:function(){var t=this;this.lastReportTimeStyleClass="blinking",setTimeout((function(){t.lastReportTimeStyleClass=""}),3500)},__init_socket_connection__:function(){var t=this;this.socketConnection=new WebSocket(this.$GCONFIG.socket_url),this.socketConnection.onopen=function(){console.log("App::socket::onopen..."),t.socketConnected=!0},this.socketConnection.onerror=function(){console.error("App::socket::error..."),t.btnSocketReconnectLoading&&(t.btnSocketReconnectLoading=!1,t.$toast.open({message:"Could not establish socket connection",type:"error"}))},this.socketConnection.onclose=function(){console.error("App::socket::close..."),t.socketConnected=!1},this.socketConnection.onmessage=function(e){console.log("App::socket::onmessage - data:",e.data);var o=JSON.parse(e.data);if("EVENT"===o.channel)t.alertSheet||t.alertCount++,o["time"]=new Date,t.$store.commit("ADD_ALERT_MESSAGE",o),"Monit"===o.service&&t.hostsReport.forEach((function(e){e.hostname===o.host&&("error"===o.type?(e.alertMessage=o.message,console.warn("Var before delete:",JSON.stringify(t.hostsNextReportTime)),delete t.hostsNextReportTime[e.id],console.warn("Var after delete:",JSON.stringify(t.hostsNextReportTime))):"success"===o.type&&(e.alertMessage=""))}));else if("HOST"===o.channel){t.lastHostsReportTime=new Date,t.__highlight_report_time_text__();for(var s=!1,r=0;r<t.hostsReport.length;r++){var a=t.hostsReport[r];if(a.id===o.id){console.log("App::socket::onmessage - found host to update"),a.poll=o.poll,a.uptime=o.uptime,a.ram=o.ram,a.cpu=o.cpu,a.services=o.services,a.goodServices=o.goodServices,a.failServices=o.failServices,a.skipServices=o.skipServices,a.alertMessage="",t.$set(t.hostsNextReportTime,o.id,(new Date).getTime()+1e3*a.poll),s=!0;break}}s||(console.log("App::socket::onmessage - new host report"),t.hostsReport.push({id:o.id,poll:o.poll,hostname:o.hostname,uptime:o.uptime,ram:o.ram,cpu:o.cpu,services:o.services,goodServices:o.goodServices,failServices:o.failServices,skipServices:o.skipServices}),t.$set(t.hostsNextReportTime,o.id,(new Date).getTime()+1e3*o.poll))}}},test:function(){var t="".concat(this.num," Lorem ipsum dolor sit amet consectetur, adipisicing elit<br/>\n        Amet nulla obcaecati laudantium assumenda autem ex qui, non quisquam alias officia?<br/>\n        Neque nulla libero corporis eos atque facilis illum aliquid inventore?");this.num%3==0?this.$toast.open({message:t,type:"error"}):this.$toast.open({message:t}),this.num++},changeDarkMode:function(){var t=!this.app_dark_mode;this.$vuetify.theme.dark=t,this.$store.commit("SET_DARK_MODE",t)},openAlertSheet:function(){this.alertCount=0,this.alertSheet=!this.alertSheet},getAlertColor:function(t){switch(t){case"info":return"info";case"success":return"success";case"error":return"error";case"warning":return"warning";default:return""}},alertReset:function(){this.alertCount=0,this.$store.commit("CLEAR_ALERT_MESSAGES")},reconnectSocket:function(){console.log("App::reconnectSocket()..."),this.btnSocketReconnectLoading=!0,this.__init_socket_connection__()},totalGoodHosts:function(){var t=this;return Object.keys(this.hostsState).filter((function(e){return"good"===t.hostsState[e]})).length},onHostStateChanged:function(t){console.log("App::onHostStateChanged() - payload:",t),this.$set(this.hostsState,t.id,t.status)}}},D=j,E=(o("89e6"),o("0798")),L=o("7496"),O=o("40dc"),N=o("5bc1"),I=o("4ca6"),$=o("e4e5"),H=o("288c"),F=o("a523"),P=o("ce7e"),G=o("553a"),B=o("f6c4"),q=o("3129"),W=o("2a7f"),J=Object(h["a"])(D,r,a,!1,null,"55711fea",null),K=J.exports;m()(J,{VAlert:E["a"],VApp:L["a"],VAppBar:O["a"],VAppBarNavIcon:N["a"],VBadge:I["a"],VBanner:$["a"],VBottomSheet:H["a"],VBtn:f["a"],VCard:_["a"],VCardText:g["b"],VCardTitle:g["c"],VCol:b["a"],VContainer:F["a"],VDivider:P["a"],VFooter:G["a"],VIcon:y["a"],VMain:B["a"],VRow:x["a"],VSkeletonLoader:q["a"],VSpacer:T["a"],VToolbarTitle:W["a"],VTooltip:V["a"]});o("d3b7");var Y=o("8c4f"),z=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",{staticClass:"home"},[s("img",{attrs:{alt:"Vue logo",src:o("cf05")}}),s("HelloWorld",{attrs:{msg:"Welcome to Your Vue.js App"}})],1)},Q=[],U=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("v-container",[s("v-row",{staticClass:"text-center"},[s("v-col",{attrs:{cols:"12"}},[s("v-img",{staticClass:"my-3",attrs:{src:o("9b19"),contain:"",height:"200"}})],1),s("v-col",{staticClass:"mb-4"},[s("h1",{staticClass:"display-2 font-weight-bold mb-3"},[t._v(" Welcome to Vuetify ")]),s("p",{staticClass:"subheading font-weight-regular"},[t._v(" For help and collaboration with other Vuetify developers, "),s("br"),t._v("please join our online "),s("a",{attrs:{href:"https://community.vuetifyjs.com",target:"_blank"}},[t._v("Discord Community")])])]),s("v-col",{staticClass:"mb-5",attrs:{cols:"12"}},[s("h2",{staticClass:"headline font-weight-bold mb-3"},[t._v(" What's next? ")]),s("v-row",{attrs:{justify:"center"}},t._l(t.whatsNext,(function(e,o){return s("a",{key:o,staticClass:"subheading mx-3",attrs:{href:e.href,target:"_blank"}},[t._v(" "+t._s(e.text)+" ")])})),0)],1),s("v-col",{staticClass:"mb-5",attrs:{cols:"12"}},[s("h2",{staticClass:"headline font-weight-bold mb-3"},[t._v(" Important Links ")]),s("v-row",{attrs:{justify:"center"}},t._l(t.importantLinks,(function(e,o){return s("a",{key:o,staticClass:"subheading mx-3",attrs:{href:e.href,target:"_blank"}},[t._v(" "+t._s(e.text)+" ")])})),0)],1),s("v-col",{staticClass:"mb-5",attrs:{cols:"12"}},[s("h2",{staticClass:"headline font-weight-bold mb-3"},[t._v(" Ecosystem ")]),s("v-row",{attrs:{justify:"center"}},t._l(t.ecosystem,(function(e,o){return s("a",{key:o,staticClass:"subheading mx-3",attrs:{href:e.href,target:"_blank"}},[t._v(" "+t._s(e.text)+" ")])})),0)],1)],1)],1)},X=[],Z={name:"HelloWorld",data:function(){return{ecosystem:[{text:"vuetify-loader",href:"https://github.com/vuetifyjs/vuetify-loader"},{text:"github",href:"https://github.com/vuetifyjs/vuetify"},{text:"awesome-vuetify",href:"https://github.com/vuetifyjs/awesome-vuetify"}],importantLinks:[{text:"Documentation",href:"https://vuetifyjs.com"},{text:"Chat",href:"https://community.vuetifyjs.com"},{text:"Made with Vuetify",href:"https://madewithvuejs.com/vuetify"},{text:"Twitter",href:"https://twitter.com/vuetifyjs"},{text:"Articles",href:"https://medium.com/vuetify"}],whatsNext:[{text:"Explore components",href:"https://vuetifyjs.com/components/api-explorer"},{text:"Select a layout",href:"https://vuetifyjs.com/getting-started/pre-made-layouts"},{text:"Frequently Asked Questions",href:"https://vuetifyjs.com/getting-started/frequently-asked-questions"}]}}},tt=Z,et=o("adda"),ot=Object(h["a"])(tt,U,X,!1,null,null,null),st=ot.exports;m()(ot,{VCol:b["a"],VContainer:F["a"],VImg:et["a"],VRow:x["a"]});var rt={name:"Home",components:{HelloWorld:st}},at=rt,nt=Object(h["a"])(at,z,Q,!1,null,null,null),it=nt.exports;s["default"].use(Y["a"]);var ct=[{path:"/",name:"Home",component:it},{path:"/about",name:"About",component:function(){return o.e("about").then(o.bind(null,"f820"))}}],lt=new Y["a"]({mode:"history",base:"/",routes:ct}),ut=lt,pt=(o("a434"),o("bfa9"));s["default"].use(M["a"]);var dt=new pt["a"]({key:"app-store",storage:window.localStorage,reducer:function(t){return{app_dark_mode:t.app_dark_mode,alert_messages:t.alert_messages}}}),ht=new M["a"].Store({plugins:[dt.plugin],state:{app_dark_mode:!1,alert_messages:[]},mutations:{SET_DARK_MODE:function(t,e){t.app_dark_mode=e},ADD_ALERT_MESSAGE:function(t,e){var o=99;t.alert_messages.splice(0,0,e),t.alert_messages.length>o&&t.alert_messages.pop()},CLEAR_ALERT_MESSAGES:function(t){t.alert_messages=[]}},actions:{},modules:{}}),vt=o("f309");s["default"].use(vt["a"]);var mt=new vt["a"]({theme:{}}),ft=o("b079"),_t=o.n(ft);o("ce8c");s["default"].config.productionTip=!1,s["default"].use(_t.a,{type:"success",position:"bottom",duration:5e3,pauseOnHover:!0}),u.a.get("/config.json").then((function(t){console.log("main() get config response:",t.data),s["default"].prototype.$GCONFIG=t.data,new s["default"]({router:ut,store:ht,vuetify:mt,render:function(t){return t(K)}}).$mount("#app")}))},"89e6":function(t,e,o){"use strict";var s=o("2162"),r=o.n(s);r.a},"9b19":function(t,e,o){t.exports=o.p+"img/logo.63a7d78d.svg"},cf05:function(t,e,o){t.exports=o.p+"img/logo.82b9c7a5.png"}});
//# sourceMappingURL=app.11b88de9.js.map