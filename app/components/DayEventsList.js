import React, { PureComponent } from 'react';
import { View, Text, StyleSheet } from 'react-native';
export default class DayTaskList extends PureComponent {
  constructor(props) {
    super(props)
    this.state = {

    }
  }
  async componentDidMount() {

  }
  render() {
    return (
      <View style={styles.container}>
        <Text>task list</Text>
      </View>
    )
  }
}


const styles = StyleSheet.create({
  container: {
    backgroundColor: '#fff',
  },
});