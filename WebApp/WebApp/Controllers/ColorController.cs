using System;
using WebApp.Models;
using WebApp.Services;
using Microsoft.AspNetCore.Mvc;
using System.Collections.Generic;


namespace WebApp.Controllers

{
    [Route("api/[controller]")]
    [ApiController]
    public class ColorController : ControllerBase
    {
        private readonly ColorService _colorService;
        public ColorController(ColorService colorService)
        {
            _colorService = colorService;
        }

        [HttpGet]
        public ActionResult<colorModel> Get() =>
            _colorService.Get();

    }
}
