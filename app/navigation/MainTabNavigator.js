import React from 'react';
import { Platform } from 'react-native';
import { createStackNavigator, createBottomTabNavigator, createMaterialTopTabNavigator } from 'react-navigation';

import TabBarIcon from '../components/TabBarIcon';
import HomeScreen from '../screens/HomeScreen';
import CalendarScreen from '../screens/CalendarScreen';
import MyScreen from '../screens/MyScreen';
import TodayScreen from '../screens/TodayScreen';

const config = Platform.select({
  web: { headerMode: 'screen' },
  default: {},
});

const HomeStack = createStackNavigator(
  {
    Home: HomeScreen,
  },
  config
);

HomeStack.navigationOptions = {
  tabBarLabel: 'Home',
  tabBarIcon: ({ focused }) => (
    <TabBarIcon
      focused={focused}
      name={
        Platform.OS === 'ios'
          ? `ios-information-circle${focused ? '' : '-outline'}`
          : 'md-information-circle'
      }
    />
  ),
};

HomeStack.path = '';

const CalendarStack = createStackNavigator(
  {
    Links: CalendarScreen,
  },
  config
);

CalendarStack.navigationOptions = {
  tabBarLabel: '记录',
  tabBarIcon: ({ focused }) => (
    <TabBarIcon focused={focused} name={Platform.OS === 'ios' ? 'ios-calendar' : 'md-calendar'} />
  ),
};

CalendarStack.path = '';

const MyStack = createStackNavigator(
  {
    Settings: MyScreen,
  },
  config
);

MyStack.navigationOptions = {
  tabBarLabel: '我的',
  tabBarIcon: ({ focused }) => (
    <TabBarIcon focused={focused} name={Platform.OS === 'ios' ? 'ios-contact' : 'md-options'} />
  ),
};

MyStack.path = '';

const TodayStack = createStackNavigator(
  {
    Settings: TodayScreen,
  },
  config
);

TodayStack.navigationOptions = {
  tabBarLabel: '今日',
  tabBarIcon: ({ focused }) => (
    <TabBarIcon
      focused={focused}
      name={Platform.OS === 'ios' ? 'ios-add-circle' : 'md-add-circle'} />
  )
};

TodayStack.path = '';

const tabNavigator = createBottomTabNavigator({
  HomeStack,
  TodayStack,
  CalendarStack,
  MyStack,
});

tabNavigator.path = '';

export default tabNavigator;
