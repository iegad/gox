using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using Sphinx.Forms;

namespace Sphinx
{
    public partial class MainFm : Form
    {
        public MainFm()
        {
            InitializeComponent();
        }

        private void miKrakenCGI_Click(object sender, EventArgs e)
        {
            KrakenFm.Instance.MdiParent = this;
            KrakenFm.Instance.WindowState = FormWindowState.Maximized;
            KrakenFm.Instance.Show();
        }

        private void miCerberusCGI_Click(object sender, EventArgs e)
        {
            CerberusFm.Instance.MdiParent = this;
            CerberusFm.Instance.WindowState = FormWindowState.Maximized;
            CerberusFm.Instance.Show();
        }
    }
}
