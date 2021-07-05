export default context => {
    const { route: { params } } = context
    console.log("params =>" + JSON.stringify( params))
    switch (params.page) {
      case 'tmp1':
      case 'tmp2':
          context.dynamicLayout = 'public'
          break;
        default:
          context.dynamicLayout = 'public'
          break;
    }
  }