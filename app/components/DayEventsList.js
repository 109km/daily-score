import React, { PureComponent } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { Card } from '@ant-design/react-native';
export default class DayTaskList extends PureComponent {
  constructor(props) {
    super(props)
    this.state = {
      list: props.list
    }
  }
  async componentDidMount() {

  }
  render() {

    const item = () => {
      return (
        <Card>
          <Card.Header
            title="This is title"
            thumbStyle={{ width: 30, height: 30 }}
            thumb="https://gw.alipayobjects.com/zos/rmsportal/MRhHctKOineMbKAZslML.jpg"
            extra="this is extra"
          />
          <Card.Body>
            <View style={{ height: 42 }}>
              <Text style={{ marginLeft: 16 }}>Card Content</Text>
            </View>
          </Card.Body>
          <Card.Footer
            content="footer content"
            extra="footer extra content"
          />
        </Card>
      )
    }

    return (
      <View style={styles.container}>






      </View>
    )
  }
}


const styles = StyleSheet.create({
  container: {
    backgroundColor: '#fff',
  },
});