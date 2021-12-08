import Stomp from "stompjs";
// ============================一般使用的变量============================
let number = 0;
let xStompClient = null;
export const dataAdapterInit = initialize;
export const calcSend = send;
// ============================  大函数体   ============================
function initialize()
{
  xStompClient = null;
  var url = "ws://" + window.location.host + "/stomp"//+ "127.0.0.1:3459/stomp" //

  var ws = new WebSocket(url);

  xStompClient = Stomp.over(ws);
  //xStompClient.debug = null;
  xStompClient.connect({}, () => connectCallBackSubscribe(xStompClient), error => reconnect(error, xStompClient));
}
// 断线重连
function reconnect(error, xStompClient)
{
  //连接失败时再次调用函数
  number++;
  xStompClient.connected = false;


  console.log(error)
  clearTimeout(setTimeout(initialize(), 1000 * 10));
  debugX("DataAdapter reconnect:" + number + " 次");
  return;
}
// ============================  订阅函数体  ============================
function connectCallBackSubscribe(xStompClient)
{
  number = 0;
  xStompClient.connected = true;
  xStompClient.subscribe("/topic/cmd", stompMessage => reflexA(stompMessage));
  xStompClient.subscribe("/application/control", stompMessage => reflexB(stompMessage));
  xStompClient.subscribe("/application/cmd", stompMessage => reflexC(stompMessage));

}
// ============================  解析函数体  ============================

function reflexA(stompMessage)
{
  stompMessage
}
function reflexB(stompMessage)
{
  stompMessage
}
function reflexC(stompMessage)
{
  stompMessage
}

// ============================  大函数体   ============================
/*** 发送数据 
 * @param {string} destination 地址,从配置文件中找
 * @param {{}} headers 头信息,没有的信息函数自动添加
 * @param {string} body 数据内容
 */
function send(destination, headers, body)
{
  // 自定义头信息
  //headers["userDefineHeaderA"] = "123";
  if (xStompClient && xStompClient.connected)
  {
    xStompClient.send(destination, headers, body);
  }
}

// ============================  其他  ============================
function debugX(text)
{
  console.log(text);
}