/*
 * Copyright 2002-2015 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package /* thymeleafsandbox.biglistflux */ io.stackdocker.k8sconsoleproject.webconsole;


import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.ui.freemarker.SpringTemplateLoader;
import org.springframework.web.reactive.result.view.freemarker.FreeMarkerConfigurer;
import org.springframework.web.reactive.result.view.freemarker.FreeMarkerViewResolver;

@Configuration
public class /* BigListFluxWebConfig */ Config {


    private ApplicationContext applicationContext;



    public /* BigListFluxWebConfig */ Config(final ApplicationContext applicationContext) {
        super();
        this.applicationContext = applicationContext;
    }




    /*
     * --------------------------------------------------------------------------
     * FREEMARKER CONFIGURATION (no autoconf for reactive FreeMarkerViewResolver)
     * --------------------------------------------------------------------------
     */


    @Bean
    public FreeMarkerConfigurer freeMarkerConfig() {
        // Note this is the reactive version of FreeMarker's configuration, so there is no auto-configuration yet.
        final FreeMarkerConfigurer freeMarkerConfigurer = new FreeMarkerConfigurer();
        freeMarkerConfigurer.setPreTemplateLoaders(new SpringTemplateLoader(this.applicationContext, "classpath:/templates/"));
        return freeMarkerConfigurer;
    }

    /*
     * ViewResolver for FreeMarker templates executing in NORMAL mode (only mode available for FreeMarker)
     * No limit to output buffer size, all data fully resolved in context.
     */
    @Bean
    public FreeMarkerViewResolver freeMarkerViewResolver() {
        final FreeMarkerViewResolver freeMarkerViewResolver = new FreeMarkerViewResolver("", ".ftl");
        freeMarkerViewResolver.setOrder(4);
        // TODO * Apparently no way to specify which views can be handled by this ViewResolver (viewNames property)
        return freeMarkerViewResolver;
    }



}
