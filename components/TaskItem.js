import React, { Component } from 'react';
import { View, StyleSheet } from 'react-native';
import { InputItem, DatePicker, List, Picker } from '@ant-design/react-native';

export default class TaskItem extends Component {
  state = {
    titleValue: "",
    startTimeValue: "",
    endTimeValue: "",
    timeSlices: 0,
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
    this.setState({
      titleValue: val
    });
  }
  onStartTimeChange = (val) => {
    this.setState({
      startTimeValue: val
    }, () => {
      this.calculateTimeUnit();
    });

  }
  onEndTimeChange = (val) => {
    this.setState({
      endTimeValue: val
    }, () => {
      this.calculateTimeUnit();
    });
  }
  onScoreChange = (val) => {
    this.setState({
      score: val
    }, () => {
      this.calculateScore(Number(val));
    });
  }
  calculateTimeUnit = () => {
    console.log('calculateTimeUnit', this.state.startTimeValue);
    if (this.state.startTimeValue && this.state.endTimeValue) {
      const timeD = (new Date(this.state.endTimeValue).getTime() - new Date(this.state.startTimeValue).getTime()) / 1000;
      console.log(timeD);
      const slices = Math.round(timeD / 1800);
      console.log(slices);
      this.setState({
        timeSlices: slices
      });
    }
  }
  calculateScore = (score) => {
    this.calculateTimeUnit();
    if (this.state.timeSlices > 0) {
      const totalScore = score * this.state.timeSlices;
      this.props.onScoreChange(totalScore);
    }
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
