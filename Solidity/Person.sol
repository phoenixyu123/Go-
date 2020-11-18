pragma solidity^0.4.0;//^表示向上兼容,如文中代表0.4.0~0.4.9都可以编译
//pragma 版本声明
//solidity 开发语言
//0.4表示主版本，最后0代表修复bug子版本


contract Person {//实际上合约的声明就相当于一个类
  //==class Person extends Contract
  uint _height;//声明属性前面最好加下划线
  uint _age;
  address _owner;//合约拥有者，address相当于uint160
  /* string _name; */

  function Person() {//类的构造函数
    _height = 180;
    _age = 29;
    _owner = msg.sender;

  }
  //如何实现set/get方法
  function setHeight(uint height) {//驼峰命名
    _height = height;
  }
  //get方法 读取_height
  function height() constant returns (uint) {//constant代表方法只读，减少gas消耗
    return _height;
  }

  function setAge(uint age) {
    _age = age;
  }
  function age() constant returns (uint) {
    return _age;
  }

  function owner() constant returns (address) {
    return  _owner;
  }

  //析构函数
  function kill() constant{
    if (_owner == msg.sender){
      selfdestruct(msg.sender);//msg代表当前的合约，sender代表合约的发起者
    }
  }

}
