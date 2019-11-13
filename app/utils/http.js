import Tools from './tools';

const REQUEST_URL = "http://172.16.46.204:8080/v1"

async function get(url, params) {
  try {

    // 简化的 url
    if (url.indexOf('http') < 0) {
      url = REQUEST_URL + url;
    }
    if (params) {
      let queryString = Tools.convertParamsToGetString(params);
      url = url + (queryString ? "?" + queryString : "");
    }

    const res = await fetch(url, {
      method: 'GET'
    });
    return res.json();
  } catch (err) {
    console.error(err)
  }

}

export default {
  get
}