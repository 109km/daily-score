import React, { Component } from 'react';
import { ScrollView, View, Text, StyleSheet } from 'react-native';
import { InputItem, DatePicker, List, Picker } from '@ant-design/react-native';
import TaskItem from '../components/TaskItem';


export default class TodayScreen extends Component {
  state = {
    totalScore: 0
  }
  onScoreChange = (val) => {
    console.log('onScoreChange', val);
    this.setState({
      totalScore: this.state.totalScore + Number(val)
    });
  }
  render() {
    return (
      <ScrollView style={styles.container}>
        <TaskItem onScoreChange={this.onScoreChange} />
        <TaskItem onScoreChange={this.onScoreChange} />
        <Text>总分{this.state.totalScore}</Text>
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
