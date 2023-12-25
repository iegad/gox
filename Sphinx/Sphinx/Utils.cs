using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Sphinx
{
    public class Utils
    {
        public static string HttpPost(string url, string req)
        {
            string rsp = "";

            try
            {
                HttpWebRequest httpReq = (HttpWebRequest)WebRequest.Create(url);
                httpReq.Method = "POST";
                httpReq.ContentType = "application/json";
                using (var sw = new StreamWriter(httpReq.GetRequestStream())) {
                    sw.Write(req);
                }

                HttpWebResponse httpRsp = (HttpWebResponse)httpReq.GetResponse();
                using (var sr = new StreamReader(httpRsp.GetResponseStream())) {
                    rsp = sr.ReadToEnd();
                }
            }
            catch(Exception ex)
            {
                MessageBox.Show(ex.Message);
            }

            return rsp;
        }

    }
}
