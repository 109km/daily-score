import React, { PureComponent } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import { Card } from '@ant-design/react-native';
export default class DayTaskList extends PureComponent {
  constructor(props) {
    super(props)
  }
  async componentDidMount() {

  }
  render() {

    const item = (data, index) => {
      return (
        <Card key={`item-card-${index}`}>
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

    const list = () => {
      return (
        <View style={styles.list}>
          {this.props.list.map(item)}
        </View>
      )
    }

    return (
      <View style={styles.container}>
        {list()}
      </View>
    )
  }
}


const styles = StyleSheet.create({
  container: {
    backgroundColor: '#fff',
  },
  list: {
    backgroundColor: '#f9f9f9'
  }
});