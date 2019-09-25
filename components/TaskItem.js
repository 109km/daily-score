import React, { Component } from 'react';
import { View, StyleSheet } from 'react-native';
import { InputItem, DatePicker, List, Picker } from '@ant-design/react-native';

export default class TaskItem extends Component {
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
      <View style={styles.container}>
        <View>
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
        </View>
      </View>
    );
  }
}
const styles = StyleSheet.create({
  container: {
    paddingTop: 15,
    marginBottom: 15,
    backgroundColor: '#fff',
  },
});
