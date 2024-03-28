import { ConfigProvider, Layout } from 'antd';

const { Header, Content, Footer, Sider } = Layout;

interface ILayoutElementStyle {
  background: string;
  backdropFilter?: string;
  height?: string;
  borderRadius?: string;
  width?: string
}

const siderStyle: ILayoutElementStyle = {
  background: "transparent",
  backdropFilter: "blur(20px)",
};

const mainLayoutStyle: ILayoutElementStyle = {
  ...siderStyle,
  height: "100vh",
};


const childLayoutStyle: ILayoutElementStyle = {
  ...siderStyle,
};

const headerStyle: ILayoutElementStyle = {
  ...siderStyle,
  height: "10vh",
};

const footerStyle: ILayoutElementStyle = {
  background: "#ffffff",
};

const contentStyle: ILayoutElementStyle = {
  ...footerStyle,
  borderRadius: "10px 10px 0 0",
};

function App(): JSX.Element {
  return (
    <ConfigProvider
      theme={{
        token: {
          // Seed Token
          borderRadius: 10,
          // bodyBg: 

          // Alias Token
          colorBgContainer: '#fff',
        },
      }}
    >
      <Layout hasSider style={mainLayoutStyle}>
        <Sider width="18%" style={siderStyle}></Sider>
        <Layout style={childLayoutStyle}>
          <Header style={headerStyle}></Header>
          <Content style={contentStyle}></Content>
          <Footer style={footerStyle}></Footer>
        </Layout>
      </Layout>
    </ConfigProvider>
  );
}

export default App;
