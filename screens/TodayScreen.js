import React, { Component } from 'react';
import { ScrollView, View, StyleSheet } from 'react-native';
import { InputItem, DatePicker, List, Picker } from '@ant-design/react-native';
import TaskItem from '../components/TaskItem';


export default class TodayScreen extends Component {
  state = {
    titleValue: "",
    startTimeValue: "",
    endTimeValue: "",
    scoresArr: [
      { label: "-1", value: -1 },
      { label: "-0.5", value: -0.5 },
      { label: "0", value: 0 },
      { label: "0.5", value: 0.5 },
      { label: "1", value: 1 }
    ],
    score: []
  }
  onTitleChange = (val) => {
    console.log(val);
    this.setState({
      titleValue: val
    });
  }
  onStartTimeChange = (val) => {
    console.log(val);
    this.setState({
      startTimeValue: val
    });
  }
  onEndTimeChange = (val) => {
    this.setState({
      endTimeValue: val
    });
  }
  onScoreChange = (val) => {
    console.log('onScoreChange', val);
    this.setState({
      score: val
    });
  }
  render() {
    return (
      <ScrollView style={styles.container}>
        <TaskItem />

        <TaskItem />
        {/* <View>
          <InputItem
            value={this.state.titleValue}
            onChange={this.onTitleChange}
            placeholder="输入时间片的内容"
          >
            干啥了
          </InputItem>
        </View>
        <View>
          <DatePicker
            value={this.state.startTimeValue}
            mode="time"
            defaultDate={new Date()}
            minuteStep={5}
            onChange={this.onStartTimeChange}
            format="HH:mm"
          >
            <List.Item arrow="horizontal">开始时间</List.Item>
          </DatePicker>
        </View>
        <View>
          <DatePicker
            value={this.state.endTimeValue}
            mode="time"
            defaultDate={new Date()}
            minuteStep={5}
            onChange={this.onEndTimeChange}
            format="HH:mm"
          >
            <List.Item arrow="horizontal">结束时间</List.Item>
          </DatePicker>
        </View>
        <View>
          <Picker
            data={this.state.scoresArr}
            cols={1}
            value={this.state.score}
            onChange={this.onScoreChange}
          >
            <List.Item arrow="horizontal">单位时间得分</List.Item>
          </Picker>
        </View> */}
      </ScrollView>
    );
  }
}

TodayScreen.navigationOptions = {
  title: '今日',
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    paddingTop: 15,
    backgroundColor: '#fff',
  },
});
