import React, { Component } from 'react';
import { ExpoConfigView } from '@expo/samples';

export default class MyScreen extends Component {
  /**
   * Go ahead and delete ExpoConfigView and replace it with your content;
   * we just wanted to give you a quick view of your config.
   */
  render() {
    return <ExpoConfigView />;
  }
}

MyScreen.navigationOptions = {
  title: '我的',
};
