using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.HttpsPolicy;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Options;
using Microsoft.Extensions.Logging;
using WebApp.Models;
using WebApp.Services;

namespace WebApp
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            services.AddCors(options =>
            {
                options.AddDefaultPolicy(
                    builder =>
                    {
                        builder.WithOrigins("*")
                                            .AllowAnyHeader()
                                            .AllowAnyMethod();
                    });
            });

            // requires using Microsoft.Extensions.Options
            services.Configure<TodolistDatabaseSettings>(
                Configuration.GetSection(nameof(TodolistDatabaseSettings)));

            services.AddSingleton<ITodolistDatabaseSettings>(sp =>
                sp.GetRequiredService<IOptions<TodolistDatabaseSettings>>().Value);

            services.AddSingleton<TodoService>();

            services.AddSingleton<ColorService>();

            services.AddControllers();
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }
            Console.WriteLine("-------------uu----------------");
            Console.WriteLine(Environment.GetEnvironmentVariable("BACKGROUND_COLOR"));
            Console.WriteLine("-------------uu----------------");
            Environment.SetEnvironmentVariable("Test1", "Value1");
            Console.WriteLine(Environment.GetEnvironmentVariable("Test1"));

            app.UseHttpsRedirection();

            app.UseRouting();

            app.UseCors();

            app.UseAuthorization();

            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
            });
        }
    }
}
