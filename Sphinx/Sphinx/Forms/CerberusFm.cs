using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace Sphinx.Forms
{
    public partial class CerberusFm : Form
    {
        static CerberusFm instance_;
        public static CerberusFm Instance
        {
            get
            {
                if (instance_ == null)
                    instance_ = new CerberusFm();
                return instance_;
            }
        }

        CerberusFm()
        {
            InitializeComponent();
        }

        public new void Close()
        {
            Dispose();
        }

        private void CerberusFm_FormClosed(object sender, FormClosedEventArgs e)
        {
            this.Close();
        }
    }
}
