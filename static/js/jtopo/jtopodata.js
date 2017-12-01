var mycanvas = {
    draw: function() {
      var ctx = document.getElementById('canvas').getContext('2d');
      var canvasContainer = document.getElementById('canvas-container');
      ctx.canvas.width = canvasContainer.offsetWidth - 2;
      ctx.canvas.height = canvasContainer.offsetHeight - 2;
    },
    initialize: function() {
      var self = mycanvas;
      self.draw();
      $(window).on('resize', function(event) {
        self.draw();
      });
    },
    CPU_ALERT: 60,
    MEM_ALERT: 75,
    FLOW_ALERT: 0.9,
    FLOW_WARN: 0.75,
    FLOW_NONE_WARN: 0.01,
    FLOW_NONE: 0,
    ALERT_NUM: 0,
    DEVICE_NUM: 0
  };
  
  getjsontop = function getjsontop(idc) {
    $.ajaxSettings.async = false;
    $.getJSON("api/jsontopo/line/" + idc, function(rtdata) {
      if (rtdata.data) {
        mycanvas.toplines = rtdata.data;
      };
    });
  };
  
  getDedicatedLines = function getDedicatedLines(lines) {
    $.ajaxSettings.async = false;
    $.getJSON("/report/getlines?expressNo=" + lines, function(rtdata) {
      if (rtdata.data) {
        var lines = rtdata.data.lines;
        if (lines.length == 0) {
          return;
        }
        mycanvas.toplines = lines;
        rtdata.data.location["idc"] = "self-dedicated";
        creatTopography(rtdata.data.location, 0.5);
        $('#layoutkeep').remove();
      };
    });
  };
  
  getmantop = function getmantop(idc) {
    $.ajaxSettings.async = false;
    /*$.getJSON("api/locationget?toponame="+idc,function(rtdata){*/
    $.getJSON("api/jsontopo/location/" + idc, function(rtdata) {
      if (rtdata.data) {
        mycanvas.mantop = new Map();
        for (var node of rtdata.data) {
          mycanvas.mantop.set(node.ip, [node.x, node.y]);
        }
      };
    });
  };
  
  setInfo = function setInfo() {
    $("#cpu").text(mycanvas.CPU_ALERT.toString() + "%");
    $("#mem").text(mycanvas.MEM_ALERT.toString() + "%");
    $("#flow-alert").text((mycanvas.FLOW_ALERT * 100).toString() + "%");
    $("#flow-warn").text((mycanvas.FLOW_WARN * 100).toString() + "%");
    $("#flow-none-warn").text((mycanvas.FLOW_NONE_WARN * 100).toString() + "%");
    $("#total_1").text(mycanvas.DEVICE_NUM);
    $("#total_2").text(mycanvas.ALERT_NUM);
  };
  
  //rgb_color like "255,255,0" 
  //font like "12px Consolas"
  setNodeFont = function setNodeFont(font_color, font) {
    if (arguments.length === 2) {
      mycanvas.nodes.forEach(function(tnd) {
        console.log(font);
        tnd.font = font;
      });
    } else {
      mycanvas.nodes.forEach(function(tnd) {
        tnd.fontColor = font_color;
      });
    }
  };
  
  getLocation = function getLocation(idc) {
    var
      js = document.createElement('script'),
      head = document.getElementsByTagName('head')[0];
    js.src = 'http://10.10.17.42:1935/api/tpos?idc=' + idc;
    head.appendChild(js);
  };
  
  refreshLocation = function refreshLocation(data) {
    creatTopography(data);
  };
  
  creatTopography = function creatTopography(data, scale) {
    // remove toobar
    var toobar = $(".jtopo_toolbar").remove();
  
    mycanvas.canvas = document.getElementById('canvas');
    mycanvas.stage = new JTopo.Stage(canvas);
    showJTopoToobar(mycanvas.stage);
    if (typeof(data['idc']) != "undefined") {
      $("#getIdc").attr("placeholder", data['idc']);
      mycanvas.name = data['idc'];
    }
    //mycanvas.stage.eagleEye.visible = true;
    mycanvas.scene = new JTopo.Scene();
    mycanvas.stage.add(mycanvas.scene);
    mycanvas.nodes = new Map();
    mycanvas.dedicateds = new Map();
    mycanvas.currentNode = null;
  
    for (var ip in data) {
      if (ip == 'idc') {
        continue;
      }
      var
        x = data[ip][0] * (mycanvas.canvas.width - 55),
        y = data[ip][1] * (mycanvas.canvas.height - 55),
        t = data[ip][2];
      nodeIdc = data[ip][3];
      if (mycanvas.mantop && mycanvas.mantop.get(ip)) {
        x = mycanvas.mantop.get(ip)[0];
        y = mycanvas.mantop.get(ip)[1];
        if (typeof(scale) != "undefined") {
          y = mycanvas.mantop.get(ip)[1] * scale;
        }
      }
      switch (t) {
        case 'subnet':
          img = '/assets/js/topo/icon/cisco/grey/c_g_cloud.svg';
          break;
        case 'specialline_switch':
          img = '/assets/js/topo/icon/h3c/h3_vpn.svg';
          break;
        case 'router':
          img = '/assets/js/topo/icon/h3c/h3_router_core.svg';
          break;
        case 'public_access_switch':
          img = '/assets/js/topo/icon/h3c/h3_switch_edge.svg';
          break;
        case 'private_access_switch':
          img = '/assets/js/topo/icon/h3c/h3_switch_edge.svg';
          break;
        case 'private_access_10G_switch':
          img = '/assets/js/topo/icon/h3c/h3_switch_edge.svg';
          break;
        case 'public_access_10G_switch':
          img = '/assets/js/topo/icon/h3c/h3_switch_edge.svg';
          break;
        case 'private_core_switch':
          img = '/assets/js/topo/icon/h3c/h3_switch_core.svg';
          break;
        case 'public_core_switch':
          img = '/assets/js/topo/icon/h3c/h3_switch_core.svg';
          break;
        case 'wireless_controller':
          img = '/assets/js/topo/icon/h3c/h3_wireless_switch.svg';
          break;
        case 'manage_converge_switch':
          img = '/assets/js/topo/icon/h3c/h3_switch_edge.svg';
          break;
        default:
          img = '/assets/js/topo/icon/cisco/blue/c_b_server.svg';
      };
  
      node = mycanvas.newNode(x, y, 5, 5, ip, img);
      node.text = ip + "\n" + nodeIdc;
      node.moreInfo = new Map();
      node.moreInfo.set('title', 'Node Info');
      node.moreInfo.set('IDC', nodeIdc);
      node.moreInfo.set('IP地址', ip);
      node.moreInfo.set('设备类型', t);
      node.moreInfo.set('CPU使用率', 'unknow');
      node.moreInfo.set('内存使用率', 'unknow');
      //node.moreInfo.set('产商','');
      //node.moreInfo.set('型号','');
  
      node.addEventListener('mouseup', function(event) {
        if (event.button == 2) {
          newMsgBox(this.moreInfo);
          handler(event);
        };
      });
      node.addEventListener('mouseout', function(event) {
        $("#contextmenu").hide();
        $("#contextmenu").remove();
      });
      mycanvas.nodes.set(ip, node);
      mycanvas.DEVICE_NUM++;
    };
  
    for (var line of mycanvas.toplines) {
      var
        srcIp = line.SrcNodeIp,
        dstIp = line.DstNodeIp;
      link = mycanvas.newLink(mycanvas.nodes.get(srcIp), mycanvas.nodes.get(dstIp));
      if (mycanvas.name == "self-dedicated") {
        var link_label = line.ExpressNo + "-" + line.SrcNodeIp + "_" + line.SrcPortName + "_" + line.DstNodeIp + "_" + line.DstPortName;
        mycanvas.dedicateds.set(link_label, link);
      };
      link.moreInfo = new Map();
      link.moreInfo.set('expressNo', '<a href="/report/dedicatedlines?subPage=monitor-status&expressNo='+line.ExpressNo+'"'+ 'target=_blank'+'>'+line.ExpressNo + '详情</a>');
      link.moreInfo.set('title', 'Link Info');
      link.moreInfo.set('源端口', srcIp + ' ' + line.SrcPortName);
      link.moreInfo.set('目的端口', dstIp + ' ' + line.DstPortName);
      link.moreInfo.set('当前入流量', '');
      link.moreInfo.set('当前出流量', '');
      link.moreInfo.set('可用带宽', '');
      link.moreInfo.set('流量分析', '<a href="javascript:void(0)" target=_blank '+ 'onclick=gotoURL("' + link_label + '")>' + 'Grafana'+'</a>');
  
      link.addEventListener('mouseover', function(event) {
          newMsgBox(this.moreInfo);
          handler(event);
      });
  
      link.mouseout(function(event) {
        $("#contextmenu").hide();
        $("#contextmenu").remove();
      });
    };
    mycanvas.stage.zoomIn();
    getCpuMemInfo();
    setInfo();
  };
  
  mycanvas.getNewLocation = function getNewLocation() {
    var redisNodes = [];
    for (var ip of mycanvas.nodes) {
      var redisNode = {
        ip: ip[0],
        idc: mycanvas.name,
        x: ip[1].x,
        y: ip[1].y
      };
      redisNodes.push(redisNode);
    }
    return redisNodes;
  };
  
  mycanvas.getNodeIPs = function getNodeIPs() {
    var nodeIps = [];
    for (var ip of mycanvas.nodes) {
      nodeIps.push(ip[0]);
    }
    return nodeIps;
  };
  
  mycanvas.newNode = function newNode(x, y, w, h, text, img) {
    this.node = new JTopo.Node(text);
    node = mycanvas.node;
    node.fontColor = "255.255.0";
    node.setLocation(x, y);
    node.setSize(w, h);
    //cloudNode.setSize(30, 26);
    node.setImage(img, true);
    mycanvas.scene.add(node);
    return node;
  };
  
  mycanvas.newLink = function newLink(nodeA, nodeZ, text, dashedPattern) {
    this.link = new JTopo.Link(nodeA, nodeZ, text);
    var link = this.link;
    link.lineWidth = 2; // 线宽
    link.bundleOffset = 60; // 折线拐角处的长度
    link.bundleGap = 15; // 线条之间的间隔
    link.textOffsetY = 3; // 文本偏移量（向下3个像素）
    //link.strokeColor = JTopo.util.randomColor(); // 线条颜色随机
    link.strokeColor = "1b,a4,df"; // 线条颜色随机
    link.dashedPattern = dashedPattern;
    mycanvas.scene.add(link);
    return link;
  };
  
  mycanvas.cpuMemAlert = function cpuMemAlert(data) {
    if (data) {
      mycanvas.ALERT_NUM = 0;
      $('#alert-body').html('');
      $('#alert-marquee').html('');
    }
    for (var flow of data["flow_data"]) {
      var tmplink = mycanvas.dedicateds.get(flow.line_label);
      if (!tmplink && typeof(tmplink) == "undefined") {
        console.log(tmplink);
        break;
      }
      var tin = flow.out;
      var tout = flow.in;
      //tmpink.strokeColor = "27,27,27";
      tmplink.moreInfo.set('当前出流量', tout.toString());
      tmplink.moreInfo.set('当前入流量', tin.toString());
      tmplink.moreInfo.set('可用带宽', flow.band.toString());
      var alertInfo = '当前进出流量：' + tin.toString() + '/' + tout.toString() + ',可用带宽' + flow.band.toString();
      if (tin + tout <= mycanvas.FLOW_NONE) {
        console.log(flow.line_label);
        tmplink.dashedPattern = 5;
        mycanvas.ALERT_NUM++;
        appendTr('traffic',flow.line_label, 'alert', alertInfo);
      } else if (tin >= mycanvas.FLOW_ALERT * flow.band || tout >= mycanvas.FLOW_ALERT * flow.band) {
        tmplink.strokeColor = "237,12,14";
        mycanvas.ALERT_NUM++;
        appendTr('traffic',flow.line_label, 'alert', alertInfo);
      } else if (tin >= mycanvas.FLOW_WARN * flow.band || tout >= mycanvas.FLOW_WARN * flow.band) {
        tmplink.strokeColor = "237,128,12";
        mycanvas.ALERT_NUM++;
        appendTr('traffic',flow.line_label, 'warning', alertInfo);
      } else if (tin >= mycanvas.FLOW_NONE_WARN * flow.band || tout >= mycanvas.FLOW_NONE_WARN * flow.band) {
        tmplink.strokeColor = "85,170,119";
      } else {
        tmplink.strokeColor = "85,170,119";
        tmplink.dashedPattern = 5;
      }
    }
  
    for (var metric of data["cpu_mem_data"]) {
      var tmpnode = mycanvas.nodes.get(metric.ip);
      tmpnode.moreInfo.set('CPU使用率', metric.cpu_utilize.toString() + "%");
      tmpnode.moreInfo.set('内存使用率', metric.mem_utilize.toString() + "%");
      tmpnode.moreInfo.set('数据刷新时间', new Date(parseInt(metric.date) * 1000).toLocaleString().substr(0, 21));
      if (parseInt(metric.cpu_utilize) >= mycanvas.CPU_ALERT) {
        tmpnode.alarm = 'warning cpu';
        tmpnode.alarmColor = '255,0,0';
        tmpnode.alarmAlpha = 0.9;
        mycanvas.ALERT_NUM++;
  
        appendTr('cpu/mem',new Map( [
              ['源端口',metric.ip],
              ['目的端口','--'],
              ['expressNo','--']
        ]), 'alert', 'CPU使用率' + metric.cpu_utilize);
      } else if (parseInt(metric.mem_utilize) >= mycanvas.MEM_ALERT) {
        tmpnode.alarm = "warrning mem";
        tmpnode.alarmColor = '230,0,0';
        tmpnode.fontColor = '0,0,0';
        tmpnode.alarmAlpha = 0.7;
        mycanvas.ALERT_NUM++;
        appendTr('cpu/mem',new Map([
              ['源端口',metric.ip],
              ['目的端口',''],
              ['expressNo','']
        ]), 'alert', '内存使用率' + metric.mem_utilize);
      }
    }
    setInfo();
  };
  
  getCpuMemInfo = function getCpuMemInfo() {
    $.post("/topology/api/cpumeninfo", {
          "ips": JSON.stringify(mycanvas.getNodeIPs()),
          "topolines": JSON.stringify(mycanvas.toplines)
        },
        function(data) {
          mycanvas.cpuMemAlert(data.data);
        }
      )
      .success(function() {
        console.log("update cpumeninfo susccess.");
      })
      .error(function() {
        console.log("update cpumeninfo error");
      });
  };
  
  function newMsgBox(moreInfo) {
    var trInfo = '';
    for (var kv of moreInfo) {
      if (kv[0] == 'title') {
        continue;
      }
      trInfo = trInfo +
        '<tr>' +
        '<td style="color:#2779aa">' + kv[0] + '</td><td>' + kv[1] + '</td>' +
        '</tr>';
    };
    var w = $('<div id="contextmenu" style="display:none;z-index:99999;">').html('' +
      '<table>' +
      '<thead><tr><th>' + moreInfo.get('title') + '</th><th>&nbsp</th></tr></thead>' +
      '<tbody>' + trInfo + '</tbody>' +
      '</table>'
    );
  
    $("body").prepend(w);
  };
  
  function handler(event) {
    $("#contextmenu").css({
      top: event.pageY,
      left: event.pageX,
    }).fadeIn(350);
  };
  
  //link_label = line.ExpressNo + "-" + line.SrcNodeIp + "_" + line.SrcPortName + "_" + line.DstNodeIp + "_" + line.DstPortName;
  function appendTr(alertType,link_label, statu, info) {
    if (statu == "warning") {
      var st = '<td class="hidden-480"> <span class="label label-warning">warning</span> </td>';
    } else if (statu == "alert") {
      var st = '<td class="hidden-480"> <span class="label label-danger">alert</span> </td>';
    }
    if (alertType == "cpu/mem") {
      var tmpLink = {moreInfo:link_label};
    } else {
      var tmpLink = mycanvas.dedicateds.get(link_label);
    }
  
    var tr = '' +
      '<tr><td>' + tmpLink.moreInfo.get('expressNo')+'</td>' +
      '<td>' + tmpLink.moreInfo.get('源端口') + '</td>' +
      '<td>' + tmpLink.moreInfo.get('目的端口') + '</td>' + st +
      '<td>' + info + '</td></tr>' +
      '';
    $('#alert-body').append(tr);
    $('#alert-marquee').append('<span>' +  tmpLink.moreInfo.get('expressNo')+": " +info + '!  </span>');
  }
  
  function gotoURL(key){
    var tmplink = mycanvas.dedicateds.get(key);
    var url="http://10.6.200.8:3000/dashboard/db/liu-liang-zhan-shi?";
    var ipifname = tmplink.moreInfo.get('源端口').split(' ');
    url += "var-source=" + ipifname[0];
    url += "&var-index="+ ipifname[1];
    url += "&var-daikuan="+tmplink.moreInfo.get("可用带宽");
    window.open(url,'newwindow','height=600,width=400,top=0,left=0,toolbar=no,menubar=no,scrollbars=no, resizable=no,location=no, status=no,depended=yes');
  }
  
  
  setInterval(getCpuMemInfo, 5 * 60 * 1000);