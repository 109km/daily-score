const Tools = {
  convertParamsToGetString: (searchParams) => {
    let exportUrl = "";
    Object.keys(searchParams).map((key) => {
      exportUrl += key + '=' + searchParams[key] + '&';
    })
    exportUrl = exportUrl.substring(0, exportUrl.length - 1)
    return exportUrl;
  }
}

export default Tools;