import React, { Component } from 'react';
import { Layout, DatePicker, Button, Table, Input, message } from 'antd';
import './App.css';

const { Header, Content } = Layout;


const columns = [
  {
    title: 'id',
    dataIndex: 'id',
    key: 'id',
  },
  {
    title: '问题1',
    dataIndex: 'q1',
    key: 'q1',
  },
  {
    title: '问题2',
    dataIndex: 'q2',
    key: 'q2',
  },
  {
    title: '问题3',
    dataIndex: 'q3',
    key: 'q3',
  },
  {
    title: '姓名',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '电话',
    dataIndex: 'phone',
    key: 'phone',
  },
  {
    title: '记录时间',
    dataIndex: 'log_time',
    key: 'log_time',
  },
]


const DefaultPagination = {
  current: 1
}

class App extends Component {

  constructor() {
    super();
    this.state = {
      isLogin: false,
      username: "",
      password: "",
      pageIndex: 0,
      limit: 10,
      loading: false,
      currentDate: "",
      tableData: [
      ],
      pagination: DefaultPagination
    }
  }

  resetPagination = () => {
    this.setState({
      pageIndex: 0,
      pagination: DefaultPagination
    });
  }

  onDateChange = async (date, datestring) => {
    this.resetPagination();
    this.updateList(datestring);
  }
  getData = async () => {

    const date = this.state.currentDate;

    let url = "http://localhost:8080/v1/ad/";
    if (window.location.href.indexOf("shulanbaobei") >= 0) {
      url = 'https://www.shulanbaobei.com/v1/ad/';
    }
    const res = await window.fetch(url + "?start=" + this.state.pageIndex + "&limit=" + this.state.limit
      + (date.length > 0 ? ("&date=" + date) : ""));
    const json = await res.json();
    return json;
  }

  loginUser = () => {
    if (this.state.username === "admin" && this.state.password === "admin@20191234") {
      this.setState({
        isLogin: true
      });
    } else {
      message.error('用户名和密码错误');
    }
  }
  onChangeUserName = (e) => {
    this.setState({
      username: e.currentTarget.value
    });
  }
  onChangePassword = (e) => {
    this.setState({
      password: e.currentTarget.value
    });
  }
  onTableChange = (pagination, filters, sorters) => {
    const newPagination = { ...this.state.pagination };
    newPagination.current = pagination.current;
    let pageIndex = (pagination.current - 1) * this.state.limit;
    this.setState({
      pagination: newPagination,
      pageIndex: pageIndex
    }, () => {
      this.updateList(this.state.currentDate);
    });

  }
  updateList = async (date = "") => {
    this.setState({
      loading: true,
      currentDate: date
    }, async () => {
      const res = await this.getData(date);
      this.setState({
        loading: false
      });
      if (res.code === 0) {
        const pagination = { ...this.state.pagination };
        pagination.total = res.data.total;
        pagination.pageSize = this.state.limit;
        const list = res.data.list;
        this.setState({
          tableData: list,
          pagination: pagination
        });
      } else {
        message.error(res.message);
      }
    });

  }
  async componentDidMount() {
    this.updateList();
  }
  render() {
    return (
      <div className="App">
        {this.state.isLogin ?
          <Layout>
            <Header>
              <DatePicker onChange={this.onDateChange} />
            </Header>
            <Content>
              <Table
                columns={columns}
                rowKey={record => record.id}
                loading={this.state.loading}
                dataSource={this.state.tableData}
                pagination={this.state.pagination}
                onChange={this.onTableChange}
              />
            </Content>
          </Layout>
          :
          <Layout>
            <Header>
              <h2 className="login-tips">登录</h2>

            </Header>
            <Content>
              <div className="login-box">
                <Input placeholder="用户名" value={this.state.username} onChange={this.onChangeUserName} />
                <Input type="password" placeholder="密码" value={this.state.password} onChange={this.onChangePassword} />
                <Button type="primary" className="btn-login" onClick={this.loginUser}>登录</Button>
              </div>
            </Content>
          </Layout>
        }

      </div>
    );
  }
}

export default App;