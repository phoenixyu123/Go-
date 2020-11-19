pragma solidity^0.4.0;

contract Fundation {

  struct funder {
    address funderAddr;
    uint256 fundMoney;//当前捐助
  }
  struct needer {
    address neederAddr;
    uint256 finalGoal;
    uint256 currentMoney;
    uint fundAmount;//捐助次数 
    mapping (uint => funder) funderInfo;
  }
  uint neederID;//存储合约发起人
  mapping (uint => needer) neederInfo;//存储合约发起人的个人信息
  //合约发起人信息需要公开


//新增needer
  function newNeeder(address _neederAddr,uint256 _finalGoal)public returns(uint,address,uint256) {
    neederID++;//序号递增
    neederInfo[neederID] = needer({
      neederAddr : msg.sender,
      finalGoal : _finalGoal,
      currentMoney : 0,
      fundAmount : 0
    });

    return (neederID,neederInfo[neederID].neederAddr,neederInfo[neederID].finalGoal);//1,0x5B38,100000000000
  }
  //zhuan'qian赚钱 
  function fundation(uint _neederID,address _funderAddr) payable returns(bool){
      
    needer storage _needer = neederInfo[_neederID];
    //   neederInfo[_neederID].currentMoney+=msg.value;
    _needer.currentMoney+=msg.value;
    _needer.fundAmount ++;
    _needer.funderInfo[_needer.fundAmount] = funder({
        funderAddr : _funderAddr,
        fundMoney : msg.value
    });
      
    return true;
  }
  //查看当前合约余额 
  function getContractBalance()view public returns(uint256){//所有的捐款都会被暂时存储到合约中 
      return this.balance;
  }
  //查看捐款状态 

  function getStatus(uint _neederID)view public returns(uint,uint256,uint256){//返回编号 、金额 、比例 
    needer memory _needer = neederInfo[_neederID];
    uint256 currentPercent = (_needer.currentMoney/_needer.finalGoal)*100;//百分比 
    return (_neederID,_needer.currentMoney,currentPercent);
  }

}
