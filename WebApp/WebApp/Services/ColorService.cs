using WebApp.Models;
using MongoDB.Driver;
using System.Collections.Generic;
using System.Linq;
using System;

namespace WebApp.Services
{
    public class ColorService
    {

        //private colorModel _color;

        public ColorService()
        {
            //colorModel _color = n
        }

        public colorModel Get() =>
            new colorModel {
                color = Environment.GetEnvironmentVariable("BACKGROUND_COLOR") == null ? "#FFFFFF" : Environment.GetEnvironmentVariable("BACKGROUND_COLOR")
            };


}
}