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

package io.stackdocker.k8sconsoleproject.webconsole.business.services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;
import io.stackdocker.k8sconsoleproject.webconsole.business.entities.Variety;
import io.stackdocker.k8sconsoleproject.webconsole.business.entities.repositories.VarietyRepository;

@Service
public class VarietyService {
    
    @Autowired
    private VarietyRepository varietyRepository;
    
    
    public VarietyService() {
        super();
    }
    
    
    
    public Flux<Variety> findAll() {
        return this.varietyRepository.findAll();
    }

    public Mono<Variety> findById(final Integer id) {
        return this.varietyRepository.findById(id);
    }
    
}
