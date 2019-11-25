import React, { Component } from 'react';
import { ScrollView, View, Text, StyleSheet } from 'react-native';
import { ExpoLinksView } from '@expo/samples';
import { Calendar } from 'react-native-calendars';
import dayjs from 'dayjs';
import DayEventsList from '../components/DayEventsList';
import http from '../utils/http';


const CalendarDefaultStyle = {
  mainColor: '#1890ff',
  unactiveColor: '#999'
}
const CurrentDateStyle = {
  selected: true,
  marked: true,
  selectedColor: CalendarDefaultStyle.mainColor
};
const CurrentDateUnactiveStyle = {
  selected: true,
  marked: true,
  selectedColor: CalendarDefaultStyle.unactiveColor
}
const SelectedDateStyle = {
  selected: true,
  selectedColor: CalendarDefaultStyle.mainColor
}

export default class CalendarScreen extends Component {
  constructor(props) {
    super(props);
    const currentDate = this.getCurrentDate();
    const maxDate = this.getNextYear();
    const markedDates = Object.create(null);
    markedDates[currentDate] = CurrentDateStyle;
    this.state = {
      markedDates: markedDates,
      currentDate: currentDate,
      maxDate: maxDate,
      todayList: [1, 2, 3],
      lastSelectedDate: ""
    }
  }
  async componentDidMount() {
    const res = await http.get('http://172.16.44.206:8080/v1/event/', {
      date: this.state.currentDate
    });
    console.log(res);
  }
  getCurrentDate = () => {
    return dayjs().format('YYYY-MM-DD');
  }
  getNextYear = () => {
    return dayjs().add(1, 'year').format('YYYY-MM-DD');
  }
  setSelectedDate = (dateString) => {
    const isToday = dateString === this.state.currentDate ? true : false;
    const markedDates = Object.assign(Object.create(null), this.state.markedDates);
    let lastSelectedDate = this.state.lastSelectedDate;

    if (!isToday) {
      markedDates[dateString] = SelectedDateStyle;
      markedDates[this.state.currentDate] = CurrentDateUnactiveStyle;
      if (lastSelectedDate.length > 0) {
        delete markedDates[lastSelectedDate];
      }
      lastSelectedDate = dateString;
    } else if (lastSelectedDate.length > 0) {
      delete markedDates[lastSelectedDate];
      markedDates[this.state.currentDate] = CurrentDateStyle;
      lastSelectedDate = "";
    }

    this.setState({
      lastSelectedDate: lastSelectedDate,
      markedDates: markedDates
    });
  }
  onDayPress = (date) => {
    const pressedDate = date.dateString;
    this.setSelectedDate(pressedDate);
  }
  render() {

    return (
      <ScrollView style={styles.container}>
        <Calendar
          current={this.state.currentDate}
          minDate={`2019-01-01`}
          maxDate={this.state.maxDate}
          onDayPress={this.onDayPress}
          markedDates={this.state.markedDates}
          arrowColor={CalendarDefaultStyle.mainColor}
        />
        <View style={styles.taskList}>
          <DayEventsList list={this.state.todayList}></DayEventsList>
        </View>
      </ScrollView>
    );
  }
}

CalendarScreen.navigationOptions = {
  title: '记录',
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    paddingTop: 15,
    backgroundColor: '#fff',
  },
  taskList: {
    marginTop: 20,
    marginBottom: 20,
    marginLeft: 20,
    marginRight: 20
  }
});
