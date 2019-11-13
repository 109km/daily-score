import Tools from './tools';

const headers = {
  "Connection": ""
}
async function get(url, params) {
  try {
    if (params) {
      let queryString = Tools.convertParamsToGetString(params);
      url = url + (queryString ? "?" + queryString : "");
    }
    const res = await fetch(url, {
      method: 'GET',
      headers: headers,
    });
    return res.json();
  } catch (err) {
    console.error(err)
  }

}

export default {
  get
}